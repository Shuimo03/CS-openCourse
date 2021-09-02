## Anatomy of a database system

### 主要内容

+ Admission Control 控制管理
+ Connection Management 连接管理

数据库管理系统大致来说分为两个部分：

+  query System
+ storage System

### 查询系统 query System

一条SQL的解析过程如下:

![SQL解析过程.JPG](https://github.com/Shuimo03/CS-openCourse/blob/main/MIT6.830Notes/LC/SQL%E8%A7%A3%E6%9E%90%E8%BF%87%E7%A8%8B.JPG?raw=true)

+ Parser----> Memory Management
+ Rewriter---> Disk Space Management
+ Planner--->  Replication Services
+ Optimizer --->  Admin Utilities
+ Executor

#### Parser

Parser部分分为两步：

+ 将SQL进行词法分析，分成一个个tokens传给Parser。
+ Parser接受Tokens，随后生成语法树。

假设有一条语句 select username from userinfo,通过语法分析器会得到4个Token，其中两个是关键字(keyword)，分别是select和from。

| 关键字 | 非关键字 | 关键字 | 非关键字 |
| :----- | :------- | :----- | :------- |
| select | username | from   | userinfo |

Parser根据token产生出语法树，类似于下图(手残见谅)

![select解析过程](https://github.com/Shuimo03/CS-openCourse/blob/main/MIT6.830Notes/LC/select.PNG)

#### Rewriter

这个阶段就会把语法分析树转换为初始查询计划，这种查询计划通常是查询的代数表达式，然后初始查询计划会被转换成一个预期所需执行时间较小的等价计划。

#### Planner

#### Optimizer

Rewriter和Planner两个部分通常称为查询优化优化器，为了选择最好的查询计划，我们需要判断：

+ 查询的

#### Executor

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