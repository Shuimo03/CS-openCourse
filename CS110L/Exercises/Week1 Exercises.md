### 目的

Week1的exercises让我们熟悉C/C++的程序分析工具，前两堂课上也是一直反复提到的静态分析和动态分析。

### part1 获取代码(Getting oriented)

官网提供了两种方式，一种是命令行的方式，还有直接下载的，具体哪种，其实都无所谓。

```
wget https://web.stanford.edu/class/cs110l/assignments/week-1-exercises/week1.zip
unzip week1.zip
```

在这次作业里面不需要写任何代码(You don’t need to write any code for this assignment)，但是需要通过分析工具来找出问题，然后写成编辑成PDF，在Gradescope上提交PDF。

这次实验我们需要安装Valgrind和最新版本的LVVM(大于等于10)，LLVM里面包含clang-tidy和sanitizers。

### 作业1 UPPERCASe

我们的第一个目标是一个简单的程序，它将一个字符串作为命令行参数，并输出大写的字符串：

```
./uppercase "hello world"
HELLO WORLD
```

这个程序将输入的字符串复制到可变缓冲区(mutable buffer)内，然后替换所有小写字母。不幸的是，编写这个程序的程序员忘记了字符串在末尾处有一个空字符串(null terminator)也就是\0。

并且保持大写字符串的缓冲区太小了。快速阅读代码来理解刚刚说的问题，然后尝试使用自动化工具(automated tools)来查找这些问题。

### 静态分析

This problem will manifest no matter what input you provide, and it’s one that an experienced programmer can easily find. How does `clang-tidy` fare?

当使用clang-tidy运行1-uppercase.c的时候

```
Error while trying to load a compilation database:
Could not auto-detect compilation database for file "1-uppercase.c"
No compilation database found in /home/cola/CS-openCourse/CS110L/Exercises/week1 or any parent directory
fixed-compilation-database: Error while opening fixed database: No such file or directory
json-compilation-database: 
```

等下来分析一下这个错误代表着什么意思，根据我们在前两课学到的知识，来推测以下为什么clang-tidy没有发现这个C文件中的问题？

### 动态分析

使用Valgrind 来分析看看是否会更好，首先我们需要确定已经编译好了程序，然后运行Valgrind：

```
make

valgrind ./1-uppercase "hello world"
==649566== Memcheck, a memory error detector
==649566== Copyright (C) 2002-2017, and GNU GPL'd, by Julian Seward et al.
==649566== Using Valgrind-3.15.0 and LibVEX; rerun with -h for copyright info
==649566== Command: ./1-uppercase hello\ world
==649566==
HELLO WORLD
==649566==
==649566== HEAP SUMMARY:
==649566==     in use at exit: 0 bytes in 0 blocks
==649566==   total heap usage: 1 allocs, 1 frees, 1,024 bytes allocated
==649566==
==649566== All heap blocks were freed -- no leaks are possible
==649566==
==649566== For lists of detected and suppressed errors, rerun with: -s
==649566== ERROR SUMMARY: 0 errors from 0 contexts (suppressed: 0 from 0)
```

Valgrind也认为这个代码没有问题，但是根据讲座中的知识，为什么Valgrind在这里也失败了？

### 通过LLVM sanitizers进行动态分析

使用sanitizers我们就得到了错误：

```
==12870==ERROR: AddressSanitizer: dynamic-stack-buffer-overflow on address 0x7ffd2aa004e4 at pc 0x0000004c32dc bp 0x7ffd2aa00420 sp 0x7ffd2aa00418
```

### 快速入门make

