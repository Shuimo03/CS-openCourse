### MapRedcude question

#### MapReduce是什么？

MapRedduce是Google在2004年发明的一款处理大数据集的算法模型，在没有MapReduce之前，Google程序员需要编程实现：

+ 统计某个关键词出现的频率，计算pageRank。
+ 对大规模数据按词频排序。
+ 对多台机器上的文件进行grep等。

这些工作不可能在一台机器上运行，所以每次编写程序都需要处理以下问题：

+ 多级并行协同。
+ 网络通信
+ 处理错误
+ 提高执行效率等等。

**而MapRedcuce是在普通机器上运行大规模并行处理程序而抽象出来的编程模型**，它可以解决以上问题并提高执行效率。

#### MapReduce解决了什么问题？

#### 什么是Map？

Map可以看成是一个统计的过程，假设输入的是一个文本，会先被分割随后每一行或者说指定的某种方式给进行统计。

#### 什么是Reduce？

Redcude则是把相同的内容聚集在一起，然后输出。

#### MapReduce工作原理？

简单说下MapReduce的工作流程，这里使用一个单词统计的例子来说明工作流程，首先读取输入文件进行分割(Split)，接着Map会读取分割的内容进行统计(KV形式)，接着会把相同的单词聚集(shuffle)聚集在一起，随后Reduce会读取shuffle中的内容并统计，最后整理并输出。

接着来说下，MapReduce的一个完整的工作流程，MapReduce由master和worker组成，MapRedcude只有一个，而worker有多个，每个worker等待master分配任务，而任务也分为两种：map任务和reduce任务。它们都是并发执行的。

master和worker之间可以通过RPC来通信。

#### 如何检查worker失效？

**master会通过心跳机制来检测worker是否失效**，如果worker没有在规定的时间内完成任务或者反馈给master，那master就会判断worker失效。所有由这个失效的 worker 完成的 Map 任务被重设为初始的空闲状态，之后这些任务就可以被安排给其他的 worker同样的，worker 失效时正在运行的 Map 或 Reduce 任务也将被重新置为空闲状态，等待重新调度。

#### 如何检测master失效？

如果有多个checkpoint的话，某个master失效就会从最后一个checkpoint启动一个master进程，如果只有一个master进程的话，就会终止MapReduce运算，接着客户根据情况来重启mapreduce。

#### master如何进行调度？

调度算法

#### RPC

