## 项目介绍

通过6.830的几个Lab，我们要实现一个简单的DBMS(database management system)，这个DBMS的名字叫做SimpleDB，在Lab1里面，我们主要实现核心模块，来访问存储在硬盘(disk)上的数据。

在将来的几个lab中，我们将会实现更多的查询进程(processing)操作，就和事务，锁，还有并发查询一样，这里可能翻译的有问题，可以看下原文。

In the lab assignments in 6.830 you will write a basic database management system called SimpleDB. For this lab, you will focus on implementing the core modules required to access stored data on disk; in future labs, you will add support for various query processing operators, as well as transactions, locking, and concurrent queries.

SimpleDB使用java编写，已经为我们提供了一些没有实现的类和接口，我们只需要去实现它们，使用JUnit编写了测试来为代码评分。

同时还提供了许多单元测试，这些单元测试不用于评分，只是检验下我们的代码有没有编写正确，但是这些测试也不是全面的，就算过了测试，也不一定能过Lab，这个比较好理解，类似于LeetCode，有时候过了样例测试，但是不一定能过这道题。

除此之外，还鼓励我们自己编写测试。

### 环境配置

参考下面链接，不多做描述，基本上就是配置好JDK和git，然后选一款适合的IDE就可以了。

https://github.com/MIT-DB-Class/simple-db-hw-2021

## Getting started

SimpleDB使用Ant build tool来编译代码和运行测试，Ant类似于make，但是构建文件是在写XML和一些配套的java代码里，大部分linux都有Ant。

使用下面的命令来运行测试单元。

```
$ cd [project-directory]
$ # run all unit tests
$ ant test
$ # run a specific unit test
$ ant runtest -Dtest=TupleTest
```

## 代码结构

+ common
+ execution
+ index
+ optimizer
+ storage
+ transaction
+ Parser.java
+ ParserExeption.java
+ SimpleDb.java

### Exercises 1

实现下面两个方法:

- src/java/simpledb/storage/TupleDesc.java
- src/java/simpledb/storage/Tuple.java

如果代码正确的话，应该可以通过TupleTest和TupleDescTest，modifyRecordId应该是失败的，因为还没有实现它。在写代码之前，先弄清楚Tuple(元组)的概念，在一个关系数据库里面，表是由行和列组成的，行可以用 **元组**表示，列可以用 **属性**表示。

Tuple可以替换成field，两个都是一个意思。

在关系中除了含有属性名所在的行以外的其他行，都可以被称为元组，弄清楚这个概念就可以动手写代码了。

TupleDesc相当于schema，schema在数据库系统中是形式语言描述的一种结构，是对象的集合，可以包括各种对象，如下：

+ 表
+ 字段
+ 关系模型
+ 视图
+ 索引
+ 包
+ 存储过程
+ 子程序
+ 队列
+ 触发器
+ 数据类型
+ 序列
+ 物化视图等等

所以它们的关系，可以这样理解，假设有一张表：

| id(int) | name(string) | age(byte) |
| ------- | ------------ | --------- |
| 1       | xiaoming     | 10        |
| 2       | xiaohong     | 11        |

一个Tuple相当于(1,xiaoming,10)，一个schema就是(id(int),name(string),age(byte))。

### TupleDesc.java

```java
/**
 * @return
 *        An iterator which iterates over all the field TDItems
 *        that are included in this TupleDesc
 * */
 
     public Iterator<TDItem> iterator() {
        // some code goes here
    }
```

这里比较简单，这个方法让我们使用Iterator遍历所有的TDitem，包括TupleDesc本身。

```java
/**
 * Create a new TupleDesc with typeAr.length fields with fields of the
 * specified types, with associated named fields.
 * 
 * @param typeAr
 *            array specifying the number of and types of fields in this
 *            TupleDesc. It must contain at least one entry.
 * @param fieldAr
 *            array specifying the names of the fields. Note that names may
 *            be null.
 */
 
     public TupleDesc(Type[] typeAr, String[] fieldAr) {
        // some code goes here
    }
```

使用typeAr的长度创建一个新的TupleDesc，字段具有特定的类型和相关的命名字段，typeAr表示的意思是数组指定的字段数量和类型，它必须包含至少一个条目。

fieldAr：数组指定字段的名字，这些名字可能会是空值(null)。

```
/**
 * @return The size (in bytes) of tuples corresponding to this TupleDesc.
 *         Note that tuples from a given TupleDesc are of a fixed size.
 */
 
     public int getSize() {
        // some code goes here
        return 0;
    }
```