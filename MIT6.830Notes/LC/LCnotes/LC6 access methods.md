## access methods

因为数据量很大，所以想使用"访问方法(access methods)"，允许我们使用最少的IO操作来获取想要的数据。

如果只是试图在一个大型数据库里面查找到一条记录，比如公司数据库里面查找特定的员工，则最少的IO操作就能派上用场。

这些访问方法是常用到的字典结构( dictionary structures)，比如hash tables, tree。专门为磁盘提供。

### 不同类型的访问方法

+ heapfiles
+ index files
  + leaves as data (primary index)
  + leaves as rids  (secondary index)
  + clustered vs unclustered 
+ types of indexes
  + hash files
  + b-trees
  + r-trees

#### 时间复杂度

### 索引类型

### hash index

## 思考问题

+ 如何在海量数据中快速查找某个数据？
+ 为什么不在每个属性上创建索引？

#### 为什么不在每个属性(列)上创建索引？

