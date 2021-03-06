## 废话

这是目前做过所有公开课里面最难的一次了，很多人都说第一个Lab简单，可能是对比后面的Lab比较简单吧，这次涉及到内容太多了，比如并发，RPC，就算给答案给代码去抄，也蛮难理解的，我的代码基本上是抄@呆呆，艾特不出来，就贴上大佬的知乎个人页吧。

(https://www.zhihu.com/people/robert.peng)

这门课想获得最大的收益应该是自己独立动手完成代码，先看reading，然后看video，再回到reading，这样去动手写代码可能能够写出来。

我之前是只看了reading，看video我会犯困，里面很多概念没了解过，所以前几次觉得论文看懂了，可以去写代码了，但是还是两眼蒙蔽。

## 准备工作

Lab1是要通过MapReduce这篇论文来实现一个简单版本的分布式MapReduce，我们将要实现worker进程来调用Map和Reduce函数，通过它们来读取或写入文件，同时coordinator进程分配任务给workers，并且还要处理失败的workers，具体工作流程和细节可以去看论文，**21版本中的master被改成了coordinator，也和论文中的master对应上了**。

worker有一个以及一个以上，它们可以并行处理。master只有一个，在现实的系统中，worker应该是运在不同的机器上，但是在lab中，我们只需要运行在一台机器上就可以了。

worker通过RPC向coordinator申请任务(task)，从一个或者多个文件读取任务的输入，然后执行任务，并将任务的输出写入一个或者多个文件。

如果有worker没有在规定的时间内完成它的task，在Lab中这个时间是10秒，那coordinator应该注意到并且把这个task给其他的worker。

Lab已经给了部分代码，在main包下面，这些代码不用去修改，但是可以去在运行官方给的例子的时候，可以去调试。

我们只需要去实现mr下面的三个文件就可以了。

+ mr/coordinator.go
+ mr/worker.go
+ mr/rpc.go

### 环境配置

因为MIT6.824都是通过Go来实现的，所以需要先安装Go和学习Go，下面给出了go的安装和学习教程:

+ 安装 https://pdos.csail.mit.edu/6.824/labs/go.html
+ 教程 https://tour.golang.org/

Git的配置不用多说了，获取初始化Lab

```
$ git clone git://g.csail.mit.edu/6.824-golabs-2021 6.824
$ cd 6.824
$ ls
Makefile src
$
```

这里有一个小例子，对应了论文中的统计，这里回顾一下这个例子，首先输入的是一个文本文件：

```
0: Deer Bear River
1: Car Car River
2: Deer Car Bear
```

master分配map任务给workers，在map阶段，则workers将文本分为一行一行，并对每一行的单词出现的次数做一次统计，map会将重复的key去除。

```
Deer Bear River
Deer 1
Bear 1
River 1

Car Car River
Car 1
Car 1
River 1

Deer Car Bear
Deer 1
Car 1
Bear 1
```

回头在看论文，Map输出的应该是一对键值对。

接着第二步，再把相同的单词放在一起，通过reduce处理最后输出。

```
Bear 1
Bear 1

Car 1
Car 1
Car 1

Deer 1
Deer 1

```

最后得到的结果如下：

```
Bear 2
Car 3
Deer 2
River 2
```

这里只是个人的理解，具体的细节还是要去看论文。

## 代码思路

整体来看就是worker向coordinator申请任务，或者coordinator分配任务给worker，然后它们之间是通过RPC通信的。所以打算先从RPC入手，注意worker是可以多个并行执行的。

先来看下map函数和reduce函数，使用wc.go作为例子：

```go
//
// The map function is called once for each file of input. The first
// argument is the name of the input file, and the second is the
// file's complete contents. You should ignore the input file name,
// and look only at the contents argument. The return value is a slice
// of key/value pairs.
//



func Map(filename string, contents string) []mr.KeyValue {
   // function to detect word separators.
   ff := func(r rune) bool { return !unicode.IsLetter(r) }

   // split contents into an array of words.
   words := strings.FieldsFunc(contents, ff)

   kva := []mr.KeyValue{}
   for _, w := range words {
      kv := mr.KeyValue{w, "1"}
      kva = append(kva, kv)
   }
   return kva
}

//
// The reduce function is called once for each key generated by the
// map tasks, with a list of all the values created for that key by
// any map task.
//
func Reduce(key string, values []string) string {
   // return the number of occurrences of this word.
   return strconv.Itoa(len(values))
}
```

map函数的意思是：每个输入的文件调用一次map，filename是输入文件的名字，contents是文件的完整内容，应该忽略文件名，只关注文件的内容，map将会返回一对键值对。

就是worker如果执行的是map任务，那worker就需要从filename中读取contents，然后从contents解析出KV键值对，最后把这个kv键值对传给用户自自定义的Map函数，最后通过Map函数来生产输出中间的kv键值对，并缓存在内存中。

这里的map和Map是有去别的，Map函数是属于用户自定义的，而map是属于mapreduce这个框架的。

mapreduce函数的意思是：map任务生成的每个键都会调用一次 reduce 函数，其中包含由任何map任务。

具体的例子已经在上面讲过了。

### rpc.go

RPC简单来说就是一台主机调用另外一台主机上的方法或者函数的通信框架，本地的调用就是LPC，因为是coordinator分配任务给worker的，所以就需要定义任务的类型，类型有两种：Map和Reduce，coordinator是一个调度进程，所以就只有一个，但是worker是多进程的，所以为了区别每个worker就需要定义它们的ID，就是用进程ID作为它们的ID好了，也方便调试，最后是任务的数量和完成的时间。

根据这些信息，我们大概也就晓得要怎么完成rpc.go了。

### worker.go

worker对于我这种菜鸡来说还是有点难度的，先来看个简单的例子和处理流程，首先是master分配给worker任务，woker要去执行任务的时候，有一个前提条件就是：它得是空闲状态，如果不是空闲状态还不能分配。

接着识别是什么类型的任务，一共就两种map和reduce，目前到这里，大体思路是有了，用伪代码来表示就是这样的：

```
if(workerStatus == leisure){
	switch master.taskType:
	case Map:
		todo map
	CASE Reude:
		todo reduce
}
```

执行完任务之后，重置woker状态为空闲，但是要注意，这里的worker是并行执行的，也就是说workerA执行map任务，还有wokerB执行reduce任务，并且每个任务都有一个时间(10秒)，如果没有在规定的时间里面完成任务，那它就不能继续执行，而是把这个任务转给其他worker做。

### coordinator.go

master是负责分配任务的，同时它还需要检测所有的worker状态，还有task的任务状态，是否完成等等，应该是三个里面最难的一个，还要面对并发问题。

## 参考链接

+ [MIT6.824-2021/lab1.md at master · LebronAl/MIT6.824-2021 (github.com)](https://github.com/LebronAl/MIT6.824-2021/blob/master/docs/lab1.md)
+ [MIT 6.824 Lab 1 - 实现 MapReduce - Robert Peng's Blog (mr-dai.github.io)](https://mr-dai.github.io/mit-6824-lab1/)
+ https://www.youtube.com/watch?v=Rz8JCS9TfOQ&t=1s

