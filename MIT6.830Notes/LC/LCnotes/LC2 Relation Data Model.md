## Relation Data Model(关系数据模型)

### 数据模型

+ Hierarchical 层次模型
+ Network 网状模型
+ Relations 关系模型

### LC2重点内容

+ Data redundancy(数据冗余)
+ Data independence(数据独立性)
  + physical data independence(programs don't embed knowledge of physical structure) 物理数据独立性
  + logical data independence((logical structure can change and legacy programs can  continue to run) 逻辑数据独立性
+ High level languages 高级语言

### 层次模型

层次模型是一种用树形数据结构来描述实体和关系的数据模型，每一个记录类型都使用节点表示，记录类型之间的联系则是用结点之间的邮箱线段来表示，每一个双亲结点可以有多个子节点，但是每一个子节点只能有一个父结点(parent)，这种结构决定了采用层次模型作为数系组织方式的层次数据库系统只能处理一对多的实体联系。

有代表性的DBMS，就是IBM推出来的IMS(Information Management System)了。

![architec](D:\CS-openCourse\MIT6.830Notes\LC\architec.gif)

### IMS Physical representation

DL/1是IMS的程序语言，

DL/1 is IMS's programming language. Programming in DL/1 is done by explicitly  iterating through the database. Basic commands are:

### IMS problems IMS的问题

+ duplication of data
+  painful low level programming interface – have to program the search algorithm
+ limited physical data independence

### 关系模型

关系模型对应的关系数据库，有一些概念需要注意，首先是关系数据库，关系数据库是由表的集合组成，每个表都有唯一的名字，表中的一行代表着一组值之间的的关系，表在关系模型中对应的是关系，表中的行对应的是元组(tuple)，属表中的列对应属性(attribute)。

#### Key properties 主要特性

+ Simple representation
+ Set-oriented programming model that doesn't require "navigation"
+ No physical data model description required(!)

#### The data model 数据模型

+ All data is represented as tables of records, or tuples
  Tables are unordered sets (no duplicates) 
+ Database is one or more tables
+ Each relation has a "schema" that describes the types of the columns/fields
+ Each field is a primitive type -- not a set or relation
+ "Keys" are used to manage the relationships between records
  + "Primary key" -- unique identifier for a particular record
  + "Foreign key" reference to a particular key in another table

#### 关系代数

a very mathy way of expressing operations  over tables，"Algebra" because it is closed -- meaning that the result of every expression is a valid  set of relations

主要操作如下:

- Projection (π(T,c1, …, cn)) -- select a subset of columns c1 .. cn 
- Selection (sel(T, pred)) -- select a subset of rows that satisfy pred 
- Cross Product (T1 x T2) -- combine two tables
- Join (T1, T2, pred) = sel(T1 x T2, pred)
- Plus various set operations (UNION, DIFFERENCE, etc)

#### SQL

#### SQL和CODASYL之间的区别

 Programming no longer navigates in N-D space -- just write a program to extract  the data they want. Simple, and elegant.

 Physical independence: Note that the relational model says nothing about the physical representation at all.  Users just interact with operators on tables (sets). Tables could be sorted, hashed,  indexed, whatever

 Programs will continue to run no matter what physical representation is chosen for the  data

\- Logical independence: (How can you change the schema without breaking legacy  programs?)

### 数据独立性

### 网状模型

