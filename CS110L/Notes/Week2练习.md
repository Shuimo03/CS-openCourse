## Ownership and structs

第二周的练习主要是两部分，首先是继续熟悉我们的Rust编程，主要是第二周学到的知识点，比如所有权和引用。还有Option和Result。

第二部分是使用面向对象的方式实现一个diff程序。

###  Ownership short-answer exercises

来看一些关于所有权的例子进行热身，下面每个例子都会有这些问题：

+ 这些例子能不能通过编译？给出具体的原因。
+ 如果不能编译，能不能将它修复，然后通过编译？

如果不能确定，可以运行它们，然后通过编译器的提示来修改它们，代价就是需要

 high-level English来解释它们可不可以工作。

第一题

```rust
fn main() {
    let mut s = String::from("hello");
    let ref1 = &s;
    let ref2 = &ref1;
    let ref3 = &ref2;
    s = String::from("goodbye");
    println!("{}", ref3.to_uppercase());
}
```

这个是运行不起来的，因为

第二题

```rust
fn drip_drop() -> &String {
    let s = String::from("hello world!");
    return &s;
}
```

不能运行

## rdiff

在这个练习中，我们将要实现一个简单版本的diff。

### Milestone 1: Reading the two files into vectors of lines

这个阶段，我们实现read_file_lines，