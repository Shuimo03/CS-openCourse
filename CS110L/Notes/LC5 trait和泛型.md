## 泛型与traits

### traits

### LC5中的一些重点:

+ 大部分OOP语言使用继承类共享函数，但是

### 代码实例

```rust
#[derive(Debug, PartialEq, Clone, Copy)]
struct Point {
    x: f64,
    y: f64
}

impl Point {
    pub fn new(x: f64, y: f64) -> Self {
        Point {x: x, y: y}
    }
}

fn main() {
    let the_origin = Point::new(0.0, 0.0);
    let mut p = the_origin; // copy semantics!
    println!("p: {:?}, the_origin: {:?}", p, the_origin);
    println!("are they equal? => {}", p == the_origin);
    p.x += 10.0;
    println!("p: {:?}, the_origin: {:?}", p, the_origin);
    println!("are they equal? => {}", p == the_origin);
}
```

这几个驱动分别表示:

+ Debug：
+ PartialEq
+ Clone
+ Copy