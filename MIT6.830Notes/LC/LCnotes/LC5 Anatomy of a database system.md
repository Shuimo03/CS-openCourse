## Anatomy of a database system

### 主要内容

+ Admission Control 控制管理
+ Connection Management 连接管理

### 查询系统 query System

一条SQL的解析过程如下:

![image-20210730094324208](D:\CS-openCourse\MIT6.830Notes\LC\SQL解析过程.JPG)

+ Parser----> Memory Management
+ Rewriter---> Disk Space Management
+ Planner--->  Replication Services
+ Optimizer --->  Admin Utilities
+ Executor

### 存储系统 storage System

## SimpleDB Overview

+ cost estimation basics
+  buffer pool

### 磁盘随机读写

## 思考问题

+ What's the "cost" of a particular plan?
+ Random I/O can be a real killer (10 million instrs/seek) . When does a disk need to seek?
+ Which do you think dominates in most database systems?
+ What if I know I will not be accessing a relation again in a query? Are some records higher value than others?
+ What is optimal buffer pool caching strategy? Always LRU?