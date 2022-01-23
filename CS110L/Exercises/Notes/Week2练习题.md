## Ownership and structs

part1的部分是检查有没有学明白所有权，给了三个例子，让我们找出所有权。并把结果卸载answer.txt中。

part2是实现一个简单版本的命令行工具(command-line utility)来比较两个文件。是一个简单版本的diff。

这里提到了两个好玩的算法：

+ Myer's algorithm
+ longest common subsequence LCS 最长子序列

原文中提到，大多数实现diff的算法都是Myer's algorithm。但是在这次练习中需要实现LCS来检查。

后续两个都是挑战。

## rdiff

在本练习中，我们将是心啊一个简单版本的diff命令行工具，用来比较两个文件，更多实现这个工具的算法是Myer's algorithm，但是在这次练习中，我们还要求你计算两个文件之间的最长子序列，并使用最长子序列(LCS)来计算文件之间的差异。

我们将程序分成不同级别(原文中是milestones)，按照顺序来实现它们。

### Milestone 1: Reading the two files into vectors of lines

第一阶段，我们编辑main.rs，实现其中的read_file_lines，读取路径中的文件到vector，然后返回这个vector。如果没有上过LC3就开始写的话，课程是建议快速过一遍LC3的课堂笔记，这里也推荐下Rust程序设计中的第九章：错误处理。

读取文件的第一件事情需要先打开文件，通过Flie

### Milestone 2: Implementing the Grid interface

打开grids.rs文件，new函数初始化足够的向量空间来存储num_rows * byn_cols元素，display函数打印Grid的内容，zero函数清零所有元素。

我们需要实现get和set函数，使用get函数的时候，如果有值就应该返回它的坐标，如果没有的话就返回None。

### Milestone 3: Implementing Longest Common Subsequence

LCS的补充已经写在下面，LCS是一个算法问题：给定两个序列，abcd和adbc，它们的最长子序列是什么？根据下面给出的定义，我们可以删掉第第一个序列中的d和第二个序列中d，那结果就是abc这个就是他的最长子序列。

When diffing two files, we want to determine which lines were added or removed between them. To do this, we need to identify the lines that are common between both files. We can frame this as an LCS problem! We have two sequences of *lines*, and we want to find the longest subsequence of lines that appears in both files; those lines are the lines that were unmodified, and the other lines are those that were added or removed.

### Milestone 4: Using LCS to construct the full diff

到这里差不多，diff就写的差不多了，这一步是用LCS来完成diff，在main.rs函数中，我们需要完成以下事情：

+ 调用写好的read_file_lines函数来读取两个文件的内容。
+ 调用lcs来获取LCS Grid
+ 实现和调用以下伪代码来打印diff。

```
* let C be the grid computed by lcs()
* let X and Y be sequences
* i and j specify a location within C that you want to look at when reading out
  the diff. (This makes more sense if you understand the LCS algorithm, but
  it's not important.) When you call this function initially, just pass
  i=len(X) and j=len(Y).
function print_diff(C, X, Y, i, j)
    if i > 0 and j > 0 and X[i-1] = Y[j-1]
        print_diff(C, X, Y, i-1, j-1)
        print "  " + X[i-1]
    else if j > 0 and (i = 0 or C[i,j-1] ≥ C[i-1,j])
        print_diff(C, X, Y, i, j-1)
        print "> " + Y[j-1]
    else if i > 0 and (j = 0 or C[i,j-1] < C[i-1,j])
        print_diff(C, X, Y, i-1, j)
        print "< " + X[i-1]
    else
        print ""
```

先实现print_diff

### 可选项：rwc

### 可选项： Conway’s Game of Life

### ? operator

?操作符一般用在Result中，

### Myer's algorithm

### 最长公共子序列 LCS

LCS是在一个序列集合中(通常为两个序列)用来查找所有序列中最长子序列的问题，子序列是指这样一个新的字符串：它是由**原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串**。

[1143. 最长公共子序列 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/longest-common-subsequence/)

举个例子：

abcde，可以删掉b和d，然后就变成ace这就是一个子序列，要保证字符顺序。
