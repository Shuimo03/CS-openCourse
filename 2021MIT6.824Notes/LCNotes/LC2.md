## Infrastructure: RPC and threads

### Go Introduction

golang的优点如下:

- good support for threads 对线程友好
- convenient RPC 更方便的RPC
- type safe 类型安全
- garbage-collected (no use after freeing problems) 垃圾回收
- threads + GC is particularly attractive!
- relatively simple 相对来说比较简单

GO的学习教程推荐使用这个: https://golang.org/doc/effective_go.html

### Threads

#### 线程的特性

+ 线程之间共享内存
+ 每个线程包括一些线程状态：
  + program counter
  + register
  + stack

#### 为什么要使用线程?

因为线程同时并发的速度很快，很适合用来构造分布式系统，冰饭

#### 替代线程的方案

#### 线程的挑战

### Web crawler web爬虫

### 爬虫遇到的挑战

### Remote Procedure Call (RPC)