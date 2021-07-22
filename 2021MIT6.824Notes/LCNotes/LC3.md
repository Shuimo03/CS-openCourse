## GFS

### The Google File System

### distributed storage 分布式存储

### distributed storage hard 分布式存储的难点

+ high performance -> shard data over many servers
+ many servers -> constant faults
+ fault tolerance -> replication
+ replication - > potential inconsistencies
+ better consistency -> low performance

### consistency 一致性

思考一个问题，假设一台单机服务器，两个进程(C1,C2)同时写入信息，这个时候，又有两个进程(C3,C4)来读取信息，那C3和C4读取到的内容是C1的还是C2的？



有可能是C1的，也有可能是C2的，但是C3和C4读到的值，是一样的。

### case study：GFS

