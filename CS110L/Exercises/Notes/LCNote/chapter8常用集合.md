## 集合分类

Rust的标准集合库实现了最常用的数据结构，它的集合类可以分为以下几类：

+ 线性结构: Vec,VecDeque,LinkedList
+ Maps: HashMap,BTreeMap
+ Sets: HashSet,BTreeSet
+ Misc: binaryHeap

### 常用的集合

+ vec: 一个挨着一个地储存一系列数量可变的值。
+ 字符串(string): 是字符的集合。
+ **哈希 map**（*hash map*）允许我们将值与一个特定的键（key）相关联。这是一个叫做 *map* 的更通用的数据结构的特定实现。

这些集合指向的数据是储存在堆上的，这意味着数据的数量不必在编译时就已知，并且还可以随着程序的运行增长或缩小。

## vector

结合官网文档给出的建议，在遇到以下情况的时候，使用vector：

+ You want to collect items up to be processed or sent elsewhere later, and don’t care about any properties of the actual values being stored.
+ You want a sequence of elements in a particular order, and will only be appending to (or near) the end.
+ You want a stack.
+ You want a resizable array.
+ You want a heap-allocated array.

vector 允许我们在一个单独的数据结构中储存多于一个的值，它在内存中彼此相邻地排列所有的值。vector 只能储存相同类型的值。

初始化vector的两种方法：

+ Vec::new();
+ vec![1,2,3,4,5] //宏定义

vector可以通过数组下标的形式获取元素，也可以通过get来获取元素，而使用 get 方法以索引作为参数来返回一个 Option<&T>。

在 vector 的结尾增加新元素时，在没有足够空间将所有所有元素依次相邻存放的情况下，可能会要求分配新内存并将老的元素拷贝到新的空间中。这时，第一个元素的引用就指向了被释放的内存。借用规则阻止程序陷入这种状况。

### 遍历vecotr



### vector扩容机制

## String 字符串

> Rust 的核心语言中只有一种字符串类型：`str`，字符串 slice，它通常以被借用的形式出现，`&str`。

String类型是标准库提供的，但是并没有写进核心语言部分，它是可增长的、可变的、有所有权的、UTF-8 编码的字符串类型。

字符串的创建有两种方式：

+ 使用 to_string方法从字符串字面值创建 String。
+ 使用String::from 函数来从字符串字面值创建String。

## 哈希表map

> `HashMap<K, V>` 类型储存了一个键类型 `K` 对应一个值类型 `V` 的映射。它通过一个 **哈希函数**（*hashing function*）来实现映射，决定如何将键和值放入内存中。

### 哈希扩容

### 哈希函数

