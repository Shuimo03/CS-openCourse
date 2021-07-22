### MapReduce

+ MapReduce是什么？
+ MapReduce如何实现？
+ 什么是Map？
+ 什么是Reduce？

MapReduce是一个编程模型，同时也是一个能够处理和生成超大数据集的算法模型，程序员先创建一个Map函数处理一个KV数据集合，然后输出处理后的KV数据集合，最后在创建一个Reduce函数用来合并所有具有相同中间key值得中间value值。

MapReduce的产生背景Google程序员为了处理海量的原始数据而实现的，比如文档抓取，Web请求日志等等，也有为了处理计算机各种类型的衍生数据，比如倒排索引，Web文档的图结构的各种表示形式等。

### MapReduce解决了什么？

+ MapReduce解决什么问题？
+ MapReduce是如何解决这些问题的？

在没有MapReduce之前，如果输入的数据量特别大，比如1TB左右，又想要在短时间内处理，或者说是能够接受的时间内完成，考虑到单机的性能，就需要将这些计算分布在其他主机上，但是这就涉及到了如何并行计算，如何分发数据，如何处理错误，这些问题综合在一起就不再是简单的输入输出处理了。

而MapReduce

### Map函数

Map是接受输入并对输入进行处理的一个函数，举个例子，输入的是一辆汽车，通过map函数，最后输出的是这辆汽车原来的零件。

### Reduce函数

如果Map是对汽车进行拆解，那Reduce就是对拆解完之后的零件，进行拼装，那能不能还原回去呢？

#### 单词统计的例子

打个比方，有一个待输入的文本，首先通过Map函数，对文本进行处理，对每个单词出现的次数进行统计，然后输出成一个KV数据集合，Reduce把Map函数产生的每一个特定的词的次数累加起来。

```
From fairest creatures we desire increase,
```

Map函数处理：

(From,1)

(fairest,1)

(creatures,1)

(we,1)

(desire,1)

(increase,1)

### MapReduce应用例子

+ Distributed Grep
+ Count of URL Access Frequency 
+ Reverse Web-Link Graph
+ Term-Vector per Host
+ Inverted Index
+ Distributed Sort

### MapReduce实现

MapReduce模型的实现有很多种，具体实现还是要看使用场景，比如一种实现是适用于小型的共享内存方法的机器，另外一种实现方式则适用于大型 NUMA 架构的多处理器的主机，而有的实 现方式更适合大型的网络连接集群。

个人推荐上手使用单机多线程实现和docker实现，成本较低。

### 执行过程

通过将Map调用的输入数据自动分割为M个数据片段集合，Map调用被分布到多台机器上执行，输入的数据片段能够在不同的机器上并行处理，使用分区函数Map调用产生的中间key值被分成R个不同分区，Reduce调用也被分不到多态机器上执行，分区数量R和分数函数都是通过用户来指定。

1. 将输入的文件分成M数据片段，每个数据片段大小一般在16MB到64MB，可以通过可选参数来控制每个数据片段大小，然后通过用户程序在机器创建大量的程序副本。
2. 程序副本由两个程序组成，master和worker，通过master分配任务，由M个Map任务和R个Reduce任务将会被分配，master将一个Map任务或者Reduce任务分配给一个空闲的worker。
3. 如果worker是被分配到了map任务，那它就会读取相关输入数据片段，从输入的数据片段中解析出KV集合，然后把KV集合传递给用户自定义的Map函数，再通过Map函数生成输出的中间KV集合缓存在内存中。
4. 缓存中的KV集合通过分区函数分成R个区域，之后周期性的写入到本地磁盘上，缓存中KV集合在本地磁盘上的存储位置将会传回给master，再由master把这些存储位置传送给Reduce worker。
5. 当Reduce woker程序接收到master程序发送过来的数据存储位置信息之后，使用RPC从Map worker所在的主机上的磁盘上读取缓存数据，当Reduce worker读取了所有的中间数据之后，通过对key进行排序，让具有相同的key值聚合在一起。因为许多不同的key值会映射到相同的reduce任务上，所有就需要进行排序，如果中间数据量太大，没有办法在内存中排序的话，就需要使用外部排序。
6. 当reduce worker程序遍历排序之后的中间数据，对每一个唯一的中间key值，reduce worker程序会将这个key值和它相关的中间value值得集合传递给用户自定义得Reduce函数，Redcude函数得输出被追加到所属分区得输出文件。
7. 当所有得Map和Reduce任务都完成之后，master唤醒用户程序，在这个时候，在用户程序里对MapReduce 调用才返回。

### Master 数据结构

### 错误容忍/容错性(Fault Tolerance)

#### worker故障

#### master故障

#### 遇到失效情况的处理机制

### 使用技巧