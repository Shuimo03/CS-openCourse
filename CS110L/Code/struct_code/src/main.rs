//整个实例必须是可变的；Rust 并不允许只将某个字段标记为可变。

#[derive(Debug)]
struct Rectangle{
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32{
        self.height * self.width
    }
    /**
     * 与字段同名的方法将被定义为只返回字段中的值，而不做其他事情。
     * 这样的方法被称为 getters,Getters 很有用，
     * 因为你可以把字段变成私有的，但方法是公共的，这样就可以把对字段的只读访问作为该类型公共 API 的一部分。
     */
    fn width(&self) -> bool{
        self.width > 0
    }

}


fn main(){

    let rect1 = Rectangle{
        width:30,
        height:50
    };

    // println!(
    //     "The area of the rectangle is {} square pixels.",
    //     area(&rect1)
    // );
        rect1.area();
}