package mr

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"sort"
	"strconv"
	"strings"
)

//
// Map functions return a slice of KeyValue.
//
type KeyValue struct {
	Key   string
	Value string
}

// for sorting by key.
type ByKey []KeyValue

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

//
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue, reducef func(string, []string) string) {

	// Your worker implementation here.

	// uncomment to send the Example RPC to the coordinator.
	// CallExample()

	workerID := strconv.Itoa(os.Getegid())
	log.Printf("Worker %s started\n", workerID)

	//向Coordinator申请Task
	var lastTaskType string
	var lastTaskIndex int
	for {

		taskArgs := TaskArgs{
			workerID,
			lastTaskType,
			lastTaskIndex,
		}
		reply := TaskReply{}
		call("Coordinator.TaskReply", &taskArgs, &reply)

		if reply.TaskType == "" {
			log.Printf("Received job finish signal from coordinator")
			break
		}

		log.Printf("Received %s task %d from coordinator", reply.TaskType, reply.TaskIndex)

		switch reply.TaskType {
		case MAP:
			//读取输入数据
			log.Printf("-----------debug Map阶段----------------\n")
			file, err := os.Open(reply.MapInputFile)
			if err != nil {
				log.Fatalf("cannot open map input file %v", reply.MapInputFile)
				log.Printf("---------不能打开输入文件--------------\n")
			}
			content, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatalf("cannot read map input file %v", reply.MapInputFile)
				log.Printf("---------不能读取输入文件--------------\n")
			}
			//传递数据至MAP函数,得到中间结果
			kva := mapf(reply.MapInputFile, string(content))
			//按Key的Hash值对中间结果进行分区
			intermediate := make(map[int][]KeyValue)
			for _, kv := range kva {
				hashed := ihash(kv.Key) % reply.ReduceTaskNum
				intermediate[hashed] = append(intermediate[hashed], kv)
			}

			//写出中间结果文件
			for i := 0; i < reply.ReduceTaskNum; i++ {
				ofile, _ := os.Create(tempMapOutFile(workerID, reply.TaskIndex, i))
				for _, kv := range intermediate[i] {
					fmt.Fprintf(ofile, "%v\t%v\n", kv.Key, kv.Value)
				}
				ofile.Close()
			}

		case REDUCE:
			log.Printf("-----------debug REDUCE阶段----------------\n")
			var lines []string
			for mi := 0; mi < reply.MapTaskNum; mi++ {
				inputFile := finalMapOutFile(mi, reply.TaskIndex)
				file, err := os.Open(inputFile)
				if err != nil {
					log.Fatalf("Failed to open map output file %s: %e", inputFile, err)
				}
				context, err := ioutil.ReadAll(file)
				if err != nil {
					log.Fatalf("Failed to read map output file %s: %e", inputFile, err)
				}
				lines = append(lines, strings.Split(string(context), "\n")...)
			}

			var kva []KeyValue
			for _, line := range lines {
				if strings.TrimSpace(line) == "" {
					continue
				}
				parts := strings.Split(line, "\t")
				kva = append(kva, KeyValue{
					Key:   parts[0],
					Value: parts[1],
				})
			}
			sort.Sort(ByKey(kva))

			ofile, _ := os.Create(tempReduceOutFile(workerID, reply.TaskIndex))

			i := 0
			for i < len(kva) {
				j := i + 1
				for j < len(kva) && kva[j].Key == kva[i].Key {
					j++
				}
				values := []string{}
				for k := i; k < j; k++ {
					values = append(values, kva[k].Value)
				}
				output := reducef(kva[i].Key, values)

				fmt.Fprintf(ofile, "%v %v\n", kva[i].Key, output)
				i = j
			}
			ofile.Close()
		}
		lastTaskType = reply.TaskType
		lastTaskIndex = reply.TaskIndex
		log.Printf("Finished %s task %d", reply.TaskType, reply.TaskIndex)
	}
	log.Printf("Worker %s exit\n", workerID)
}

//
// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
//
//func CallExample() {
//
//	// declare an argument structure.
//	args := ExampleArgs{}
//
//	// fill in the argument(s).
//	args.X = 99
//
//	// declare a reply structure.
//	reply := ExampleReply{}
//
//	// send the RPC request, wait for the reply.
//	call("Coordinator.Example", &args, &reply)
//
//	// reply.Y should be 100.
//	fmt.Printf("reply.Y %v\n", reply.Y)
//}

//
// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
