fn max_num<T>(list: &[T]) -> T{
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}


fn main() {
    let number_list = vec![34, 50, 25, 100, 65];
    let res = max_num(&number_list);
    println!("The largest number is {}",res);
}

