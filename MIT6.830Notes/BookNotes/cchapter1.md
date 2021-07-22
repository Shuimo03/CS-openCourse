## overview of databases

### questions

+ What is a DBMS, in particular, a relational DBMS? 什么是数据库？什么是关系数据库？
+ Why should we consider a DBMS to manage data? 为什么要考虑使用DBMS来管理数据？
+ How is application data represented in a DBMS? DBMS中的数据如何表示？
+ How does a DBMS support concurrent access and protect data during
  system failures? DBMS如何在系统崩溃的时候，同时支持访问和保护数据？
+ What are the main components of a DBMS? DBMS的主要构成是什么？
+ Who is involved with databases in real life? 在现实生活中，谁参入数据库？

这些问题分别可以在以下章节中找到答案：

### key concept(重要概念):

+ database management
+ data independence
+ database design
+ data model
+ relational databases and queries
+ schemas
+ levels of abstraction
+ transactions
+ concurrency and locking
+ recovery and logging
+ DBMS architecture
+ database administrator
+ application pro-grammer
+ end user

### database definition and database management system

数据库是一个数据的集合，通常描述一个或者多个相关组织，而database management system(DBMS)，则是一个软件。

### managing data

### data model

+ hierarchical data model
+ relational data model
+ semantic data model
+ network model
+ object-oriented model

### relational data model

根据数据模型对数据进行的描述称为模式(schema)，在关系模型中，

### file system vs DBMS

### 数据库如何存储视频？

### ADVANTAGES of A DBMS DBMS的优势

+ Data Independence
+ Efficient Data Access
+ Data Integrity and Security
+ Data Administration
+ Concurrent Access and Crash Recovery
+ Reduced Application Development Time

### Shortcoming of DBMS DBMS的缺点

### DESCRIBING AND STORING DATA IN A DBMS

### Levels of Abstraction in a DBMS

+ conceptual schema or logical schema
+ physical schema
+ external schema

### tuple

### Data Independence

### QUERIES IN A DBMS

### relational algebra

### DDL

### DML

### SQL

### Points to Note

关于DBMS支持并发控制和恢复，有三点注意：

+ Every object that is read or written by a transaction is first locked in shared
  or exclusive mode, respectively
+ For efficient log maintenance, the DBMS must be able to selectively force
  a collection of pages in main memory to disk.
+ Periodic checkpointing can reduce the time needed to recover from a crash

### STRUCTURE OF DBMS