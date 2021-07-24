## Problem Set 1: SQL

### 介绍

这次作业主要是了解SQL编程，使用的数据库为sqlite3.0，官方推荐的教程是http://sqlzoo.net。

### 数据集(data set)

这次作业使用的数据集是现代奥林匹克运动会的120年历史资料，数据集地址：[120 years of Olympic history: athletes and results | Kaggle](https://www.kaggle.com/heesoo37/120-years-of-olympic-history-athletes-and-results)

数据集表包括以下内容：

+ athletes：远动员信息表
+ host cities：主场城市
+ regions：区域表
+ athlete events：运动员项目表

### 使用这个数据集

### 问题1 (5分)

Find all athletes who participated in an Olympic Games before they were 12. Print the id, name, and age of the athlete in the increasing order of id. Display the output like ‘1234|Miraitowa|11’.

查找所有在12岁之前参加了奥运会的运动员，输出这些运动员的id，name和age。

id是递增顺序的，输出结果如下：

1234 | Miraitowa | 11

#### 解题思路

这题应该是属于热身的，按照它说的步骤在athletes查找，最后对查找的内容进行排序就可以了。

### 问题2 (5分)

Print the medals table for “2016 Summer” Olympic Games. Print the NOC of the country, number of gold, silver, bronze, and total medals won by each country. Sort them in the decreasing order of gold, silver, and then bronze medals. Use NOC of the country in the increasing order as tie-breaker. Print rows only for countries that participated in the 2016 Summer games. You might find CASE expression useful. Display like ‘USA|10|20|30|60’

打印2016年夏季奥运会奖牌表，需要打印以下内容：

+ NOC of the country 获奖国家 
+  number of gold,silver, bronze 奖牌数量(金牌/银牌/铜牌)
+ total medals won by each country 每个国家获得的奖牌总数

按照金牌、银牌、铜牌的顺序排列，如下所示：

‘USA|10|20|30|60’

#### 解题思路

题目中提示我们可以使用CASE这个关键字，CASE相当于 if then else，这些信息都在athlete_events这张表里面，我们注意一下，题目的意思只要2016年夏季的信息，结果是要显示 获奖的国家 金牌 银牌 铜牌 总牌数量，题目中只是给了一个例子，我们这里用数据库里面真实的例子来看下。

