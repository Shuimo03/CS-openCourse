### week1

第一周的练习主要是熟悉下rust的基本语法以及一些常用的数据结构使用：

+ vector
+ hashset

第一题的三个小作业简单，就是对vector做操作：

+ 第一题：将vector的所有元素相加。
+ 第二题：去除vector中的元素，然后将元素相加。
+ 第三题：使用hashset对vector去重，然后相加。

第二题就比较麻烦了，其实就是用到的东西多一点，这里先贴上代码：

```rust
extern crate rand;
use rand::Rng;
use std::fs;
use std::io;
use std::io::Write;

const NUM_INCORRECT_GUESSES: u32 = 5;
const WORDS_PATH: &str = "words.txt";

fn pick_a_random_word() -> String {
    let file_string = fs::read_to_string(WORDS_PATH).expect("Unable to read file.");
    let words: Vec<&str> = file_string.split('\n').collect();
    String::from(words[rand::thread_rng().gen_range(0, words.len())].trim())
}

fn main() {
    let secret_word = pick_a_random_word();
    // Note: given what you know about Rust so far, it's easier to pull characters out of a
    // vector than it is to pull them out of a string. You can get the ith character of
    // secret_word by doing secret_word_chars[i].
    let secret_word_chars: Vec<char> = secret_word.chars().collect();
    // Uncomment for debugging:
    // println!("random word: {}", secret_word);

    // Your code here! :)
    
    
}
```

和rust官方的例子很像，都是猜字游戏，只不过这次是读取txt文件的。先来分析下使用了那些库：

+ extern crate rand 是使用cargo.toml中的依赖，引入自己写的或者依赖中的库。
+ use的主要用途有三种：
  + 用来引用于某个外部模块。
  + 直接使用枚举值，而无需手动加上作用域。
  + 为某个作用域下的方法或作用域创建别名

