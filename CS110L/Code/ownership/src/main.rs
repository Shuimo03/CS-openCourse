fn main() {
    let s1 = String::from("hello");

    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}

fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();
    for(i,&item) in btyes.iter().enumerate(){
        if item == b' '{
            return i;
        }
    }
    s.len();
}
