use std::fs;

fn main() {
    let contents = fs::read_to_string("input.txt")
        .expect("Something went wrong reading the file");

    for i in 4..contents.len() {
        let buffer = &contents[i-4..i];
        if is_unique(buffer) {
            println!("{} is unique and starts at {}", buffer, i);
            break;
        }
    }

    for i in 14..contents.len() {
        let buffer = &contents[i-14..i];
        if is_unique(buffer) {
            println!("{} is unique and starts at {}", buffer, i);
            break;
        }
    }

}

fn is_unique(buffer: &str) -> bool {
    let mut chars: Vec<char> = buffer.chars().collect();
    chars.sort();
    chars.dedup();
    chars.len() == buffer.len()
}
