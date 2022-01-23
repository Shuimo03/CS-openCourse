### 堆栈分配

+ 栈溢出
+ 缓冲区
+ 缓冲区溢出

### 缓冲区溢出

缓冲区溢出是一块连续的计算机内存区域，可以保存相同的数据类型，缓冲区可以是堆栈(自动变量)、堆(动态内存)和静态数据区(全局或静态)。

在C/C++中，通常使用字符数组和malloc/new之类内存分配函数实现缓冲区。溢出指数据被添加到分配给该缓冲区的内存块之外。

### Spot the overflow

### C中字符串缓冲溢出

### 类型系统 type system

### dynamic analysis 动态分析

+ 运行程序，观察程序发生了什么，并且查找有问题的行为。
+ 可以发现问题，但是只有在程序用于测试的输入上出现问题时才可以发现。
+ 一般来说是结合使用许多不同的测试输入(比如fuzzing)来运行程序，但是这个方法依然不能保证代码没有漏洞。

### static  analysis 静态分析

+ 阅读源码(source code)并找出有问题的部分。
+ 如果再简单的情况下就会很容易，比如出现C中出现的gets()。
+ 但是一般情况下不太可能。(Halting Problem)

静态分析就是在程序运行之前，在编译时完成全部的检测。

###  Halting Problem(停机问题)




### Morris Worm

### question

+ 如何通过静态分析来判断一个程序是否停止？或者退出？
+ Are there ways that we can make static analysis more tractable/helpful?
+ How do we find and prevent common mistakes in C/C++?
+ How does Rust’s type system prevent common memory safety errors?
+ How do you architect good code?
+ 如何保证一个语言是可靠的？

## 课程项目

+ Mini GDB
+ High-performance web server 高性能服务器

## 扩展

### programming language analysis

PL的分类:

+ 理论部分
+ 环境
+ 应用

语言核心分成三种:

+ 命令式语言(Java/C++/C...)
+ 函数式语言
+ 逻辑语言

不过有些函数式语言是和命令语言兼容的，比如python,JavaScript。

### 为什么需要静态分析

静态分析可以保证以下特性

+ 程序的可靠性
+ 程序的安全性
+ 编译优化
+ 程序理解
  + IDE 提示