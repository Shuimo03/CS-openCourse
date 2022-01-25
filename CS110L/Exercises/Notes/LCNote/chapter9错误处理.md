## 错误处理

Rust 将错误组合成两个主要类别：

+ 可恢复错误：可恢复错误通常代表向用户报告错误和重试操作是合理的情况，比如未找到文件。
+ 不可恢复错误：不可恢复错误通常是 bug 的同义词，比如尝试访问超过数组结尾的位置。

Rust并不像其他语言一样有异常处理，但是，有针对两种错误的方法：

+ 可恢复错误：Result<T,E>
+ 不可恢复错误：panic!

### panic!与不可恢复错误

### Result与可恢复错误

## 传播错误

### unwrap

unwrap可以帮我们简化代码，省去match部分，如果Result的值是成员Ok，unwarp会返回Ok中的值，如果是Err，unwarp就会调用panic!。

### expect

expect和unwarp类似，expect返回文件句柄或者调用panic!宏。`expect` 在调用 `panic!` 时使用的错误信息将是我们传递给 `expect` 的参数，而不像 `unwrap` 那样使用默认的 `panic!` 信息。

## ?运算符

对比下两段代码：

```rust
fn read_username_from_file() -> Result<String, io::Error>{
    let f = File::open("hello.txt");

    let mut f = match f{
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut = String::new();

    match f.read_to_string(&mut s){
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}
```

使用?运算符后:

```rust
fn read_username_from_file() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}
```

两段代码的做的事情一样，如果 `Result` 的值是 `Ok`，这个表达式将会返回 `Ok` 中的值而程序将继续执行。如果值是 `Err`，`Err` 中的值将作为整个函数的返回值，就好像使用了 `return` 关键字一样，这样错误值就被传播给了调用者。

## 什么时候使用panic!

> 