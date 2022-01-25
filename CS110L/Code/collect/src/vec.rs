#[warn(dead_code)]
pub fn print_vec(){
    let v = vec![1, 2, 3, 4, 5];

    let third: &i32 = &v[2];
    println!("The third element is {}", third);
    
    //使用 get 方法以索引作为参数来返回一个 Option<&T>
    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element."),
    }
}

enum SpreadsheetCell  {
    Int(32),
    Float(f64),
    Text(String),
}

#[warn(dead_code)]
pub fn print_foreach_vec(){
    let v = vec![100, 32, 57];
    for i in &v{
        println!("{}", i);
    }
}

pub fn print_enum_vec(){
    SpreadsheetCell::Int(3),
    SpreadsheetCell::Text(String::from("blue")),
    SpreadsheetCell::Float(10.12),
};