## LC1  Introduction

### Distributed System

一个分布式系统就是一组计算机系统在一起工作，对于终端用户来说，就像是一台机器，这一组一起工作的计算机拥有共享状态，它们同时运行，如果其中一台机器坏了，也不会影响到整个系统的正常运行。

mapReduce就是一个很好的例子，它由两个部分组成：

+ client
+ master

master分配任务，client通过RPC向master申请任务，最后处理任务，master只有一个，而client有很多个。

### build an Distributed System

- to increase capacity via parallelism
- to tolerate faults via replication
- to place computing physically close to external entities
- to achieve security via isolation

### Distributed System Challenges

### 课程Lab

+ MapReduce 实现一个简单版本的MapReduce
+ replication for fault-tolerance using Raft 实现Raft算法
+ fault-tolerant key/value store 通过Raft算法来实现一个可以容错的KV服务
+ sharded key/value store 将第三个实验分发到一系列独立的集群中，并通过运行这些独立的副本集群进行加速

### course Topics

这是这门课的重点

+ Storage
+ Communication
+ Computation

### fault tolerance 容错

### consistency  一致性

### performance 性能

## MapReduce

