### Google File System简介

+ GFS是什么？
+ GFS解决了什么问题？
+ GFS和其他分布式文件系统的相同处与不同处？

GFS是一个可扩展的分布式文件系统，用于大型分布式数据密集型应用程序，提供了运行在廉价的硬件上运行时的容错性，并且也提供了对大量客户端的搞聚合性。

### 分布式文件系统的特点

+ 性能(performance)
+ 可扩展性(scalability)
+ 可靠性(reliability)
+ 实用性(availability)

### GFS特点

GFS容忍机器出现故障，因为这是一种常态而不是一种异常，**文件系统由数百台甚至上千台廉价的机器构成**，可以被大量的客户端机器访问。

因为这些机器都是普通且廉价的，所以某些机器在任何给定的时间都不能正常工作，而某些机器将无法从当前的故障中恢复。

除了机器问题之外，还有一些是由外部原因导致的：

+ 应用bug(application bugs)
+ 操作系统bug(operating system bugs)
+ 人为失误(human errors)
+ 硬件出错(the failures of disks, memory, connectors, networking)

为了解决以上等问题，所以GFS新增了持续监控，错误检测，容错和自动恢复。

### 文件大小和IO参数

### 文件读写

大部分文件都是通过新加数据而不是覆盖原来的数据来改变的，一旦写入，文件就只能读取，而且只能按顺序读取。

### GFS原子性与同步

### GFS设计

### GFS提供的接口

GFS并没有实现标准的API，比如POSIX，实现了一些通用的API：

+ create file
+ delete file
+ ope file
+ close file
+ read file
+ write file

除此之外，GFS还有快照(snapshot)和记录追加(record)操作，快照以低成本创建文本或者目录树副本，追加操作允许多个客户端并发的将数据追加到同一个文件里面，同时保证每个客户端追加的原子性。

### GFS结构

![GFS Architecture.png](https://github.com/Shuimo03/CS-openCourse/blob/main/2021MIT6.824Notes/paperNotes/GFS%20Architecture.png?raw=true)

### 块大小(Chunk Size)

### 元信息(Metadata)

master主要存储三类元信息类型：

+ 文件和块名称(the file and  chunk namespaces)
+ 文件到块的映射(the mapping from files to chunks)
+ 每块副本的存储位置(the locations of each chunk's replicas)

### 一致性模型



