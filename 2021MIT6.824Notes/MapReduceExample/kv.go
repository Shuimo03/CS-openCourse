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
		call("Coordinator.ApplyForTask", &taskArgs, &reply)

		switch reply.taskType {
		case "":
			log.Printf("Received job finish signal from coordinator")
			break

			log.Printf("Received %s task %d from coordinator", reply.taskType, reply.taskIndex)
		case MAP:
			//读取输入数据
			file, err := os.Open(reply.mapInputFile)
			if err != nil {
				log.Fatalf("cannot open map input file %v", reply.mapInputFile)
			}
			content, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatalf("cannot read map input file %v", reply.mapInputFile)
			}
			//传递数据至MAP函数,得到中间结果
			kva := mapf(reply.mapInputFile, string(content))
			//按Key的Hash值对中间结果进行分区
			intermediate := make(map[int][]KeyValue)
			for _, kv := range kva {
				hashed := ihash(kv.Key) % reply.reduceTaskNum
				intermediate[hashed] = append(intermediate[hashed], kv)
			}

			//写出中间结果文件
			for i := 0; i < reply.reduceTaskNum; i++ {
				ofile, _ := os.Create(tempMapOutFile(workerID, reply.taskIndex, i))
				for _, kv := range intermediate[i] {
					fmt.Fprintf(ofile, "%v\t%v\n", kv.Key, kv.Value)
				}
				ofile.Close()
			}

		case REDUCE:
			var lines []string
			for mi := 0; mi < reply.mapTaskNum; mi++ {
				inputFile := finalMapOutFile(mi, reply.taskIndex)
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

			ofile, _ := os.Create(tempReduceOutFile(workerID, reply.taskIndex))

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
		lastTaskType = reply.taskType
		lastTaskIndex = reply.taskIndex
		log.Printf("Finished %s task %d", reply.taskType, reply.taskIndex)
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
