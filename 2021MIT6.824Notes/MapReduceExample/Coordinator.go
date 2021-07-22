package mr

import (
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type Coordinator struct {
	// Your definitions here.
	lock           sync.Mutex //保护共享信息,避免并发冲突
	taskStatus     string     // 任务状态
	numMap         int
	numReduce      int
	tasks          map[string]Task
	availableTasks chan Task
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
//func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
//	reply.Y = args.X + 1
//	return nil
//}

func (c *Coordinator) TaskReply(args *TaskArgs, reply *TaskReply) error {

	//记录上一个worker的task已经完成了。
	if args.LastTaskType != "" {
		c.lock.Lock()
		lastTaskID := GenTaskID(args.LastTaskType, args.LastTaskIndex)

		if task, exists := c.tasks[lastTaskID]; exists && task.WorkerID == args.WorkerID {
			log.Printf(
				"Mark %s task %d as finished on worker %s\n",
				task.TaskType, task.Index, task.WorkerID)

			switch args.LastTaskType {
			case MAP:
				for ri := 0; ri < c.numReduce; ri++ {
					err := os.Rename(
						tempMapOutFile(args.WorkerID, args.LastTaskIndex, ri),
						finalMapOutFile(args.LastTaskIndex, ri))

					if err != nil {
						log.Fatalf("Failed to mark map output file `%s` as final: %e", tempMapOutFile(args.WorkerID, args.LastTaskIndex, ri), err)
					}
				}

			case REDUCE:
				err := os.Rename(
					tempReduceOutFile(args.WorkerID, args.LastTaskIndex),
					finalReduceOutFile(args.LastTaskIndex))

				if err != nil {
					log.Fatalf("Failed to mark reduce output file `%s` as final: %e",
						tempReduceOutFile(args.WorkerID, args.LastTaskIndex), err)
				}
				delete(c.tasks, lastTaskID)

				// 当前阶段所有的任务都完成了,进入下一个阶段
				if len(c.tasks) == 0 {
					c.transit()
				}
			}
		}
		c.lock.Unlock()
	}

	//获取一个可用task并返回
	task, ok := <-c.availableTasks
	if !ok {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	log.Printf("Assign %s task %d to worker %s\n", task.TaskType, task.Index, args.WorkerID)

	task.WorkerID = args.WorkerID
	task.Deadline = time.Now().Add(10 * time.Second)
	c.tasks[GenTaskID(task.TaskType, task.Index)] = task
	reply.TaskType = task.TaskType
	reply.TaskIndex = task.Index
	reply.MapInputFile = task.MapInputFile
	reply.MapTaskNum = c.numMap
	reply.ReduceTaskNum = c.numReduce
	return nil
}

func (c *Coordinator) transit() {
	switch c.taskStatus {
	case MAP:
		log.Printf("All MAP tasks finished. Transit to REDUCE stage\n")
		c.taskStatus = REDUCE

		// 生成Reduce Task
		for i := 0; i < c.numReduce; i++ {
			task := Task{
				TaskType: REDUCE,
				Index:    i,
			}
			c.tasks[GenTaskID(task.TaskType, task.Index)] = task
			c.availableTasks <- task
		}

	case REDUCE:
		log.Printf("All REDUCE tasks finished. Prepare to exit\n")
		close(c.availableTasks)
		c.taskStatus = "" //使用空字符串表示该作业完成
	}
}

//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	// Your code here.
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.taskStatus == ""

}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		taskStatus:     MAP,
		numMap:         len(files),
		numReduce:      nReduce,
		tasks:          make(map[string]Task),
		availableTasks: make(chan Task, int(math.Max(float64(len(files)), float64(nReduce)))),
	}

	for i, file := range files {
		task := Task{
			TaskType:     MAP,
			Index:        i,
			MapInputFile: file,
		}
		c.tasks[GenTaskID(task.TaskType, task.Index)] = task
		c.availableTasks <- task
	}

	c.server()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)

			c.lock.Lock()

			for _, task := range c.tasks {
				if task.WorkerID != "" && time.Now().After(task.Deadline) {

					//回收并重新分配
					log.Printf(
						"Found timed-out %s task %d previously running on worker %s. Prepare to re-assign",
						task.TaskType, task.Index, task.WorkerID)
					task.WorkerID = ""
					c.availableTasks <- task
				}
			}
			c.lock.Unlock()
		}
	}()

	return &c
}

func GenTaskID(TaskType string, index int) string {
	return fmt.Sprintf("%s-%d", TaskType, index)
}
