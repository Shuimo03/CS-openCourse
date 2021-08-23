### Introduction

Lab2一共分为四个部分，整个Lab是实现一个具有容错性的kv存储系统，在Lab 2A中是实现Raft，一种复制状态协议(replicated state machine protocol)，接下来三个部分的lab都是基于raft实现的。

复制服务通过在多个副本服务器上存储其状态（即数据）的完整副本来实现容错，复制机制(Replication)运行服务器在遇到一些问题的情况下继续运行，比如崩溃或网络不稳定，所以遇到的挑战是故障可能会导致副本保存着不同的数据副本。

