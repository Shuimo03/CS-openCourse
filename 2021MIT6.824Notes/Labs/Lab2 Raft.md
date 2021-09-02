### Introduction

Lab2一共分为四个部分(2A-2D)，整**个Lab是实现一个具有容错性的kv存储系统**，在Lab 2A中是实现Raft，一种复制状态协议(replicated state machine protocol)，接下来三个部分的lab都是基于raft实现的。

复制服务通过在多个副本服务器上存储其状态（即数据）的完整副本来实现容错，复制机制(Replication)运行服务器在遇到一些问题的情况下继续运行，比如崩溃或网络不稳定，所以遇到的挑战是故障可能会导致副本保存着不同的数据副本。

Raft把客户端请求组织成一个序列，这个序列称为日志，并且确保所有副本服务器能够看到相同的日志，每个副本按照日志顺序执行客户端请求，将它们应用到服务器状态的本地副本中。

### Lab2A

2A主要是实现leader election和心跳机制，2A的目标是选出一个单独的leader ，如果没有失败则该leader还是leader，

### 准备工作

2A的工作是建立在Lab1之上的，主要是在raft.go里面实现，

Lab2A需要实现的内容如下：

+ RequestVoteArgs
+ RequestVote
+ RequestVoteReply
+ AppendEntries
+ type Raft struct{}

### RequestVoteArgs

修改Make()结构来创建一个后台的goroutine，当它有一段时间没有接收到其他对等方消息的时候，就会通过发送RequestVote RPC来启动leader选举。
