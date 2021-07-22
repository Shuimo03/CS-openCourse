## Review

### 关系代数和SQL之间的转换

- SELECT ==> Projection list 
- FROM ==> all tables referenced
-  JOIN ==> join expressions
-  WHERE ==> select expressions

### 关系代数和SQL之间的区别

### Recap: SQL vs CODASYL 

#### Programming no longer navigates in N-D space

#### Physical independence

Note that the relational model says nothing about the physical representation at all. Users just interact with operators on tables. Tables could be sorted, hashed, indexed, whatever

关系模型并没有显示关于物理表示，用户只是和表进行操作，可以对表进行排序，散列，索引等操作。

####  Logical independence

### 视图

### Schema Normalization 模式规范化

目标：生产一组没有冗余的Schema。

#### 为什么要消除冗余？

+ Because it leads to various anomalies when you try to update a table.
+  Because it wastes space

#### 通过什么来消除冗余？

### functional dependency (FD)

+ FD1
+ FD2
+ FD3

