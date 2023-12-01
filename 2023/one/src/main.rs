use std::fs;

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Should have been able to read the file");

    let digit_lines = contents
        .lines()
        .map(|line| get_digits(line))
        .collect::<Vec<String>>();

    let mut sum = 0;

    for line in digit_lines {
        if line.len() == 0 {
            continue;
        } else {
            let combined: String;
            let first = line.chars().nth(0).unwrap();
            if line.len() == 0 {
                combined = format!("{}{}", first, first);
            } else {
                let last = line.chars().nth(line.len() - 1).unwrap();
                combined = format!("{}{}", first, last);
            }
            sum += combined.parse::<i32>().unwrap();
        }
    }

    println!("Sum: {}", sum);
}

fn get_digits(line: &str) -> String {
    let mut digits = String::new();

    let vec = line.chars().collect::<Vec<char>>();
    for i in 0..vec.len() {
        let current = vec[i];
        let remaining_chars = vec.len() - i;

        if current.is_digit(10) {
            digits.push(current);
        }

        if current == 'o' {
            if remaining_chars < 3 {
                continue;
            }
            if vec[i + 1] == 'n' && vec[i + 2] == 'e' {
                digits.push('1');
                continue;
            }
        }

        if current == 't' {
            if remaining_chars < 3 {
                continue;
            }
            if vec[i + 1] == 'w' && vec[i + 2] == 'o' {
                digits.push('2');
                continue;
            }
        }

        if current == 't' {
            if remaining_chars < 5 {
                continue;
            }
            if vec[i + 1] == 'h' && vec[i + 2] == 'r' && vec[i + 3] == 'e' && vec[i + 4] == 'e' {
                digits.push('3');
                continue;
            }
        }

        if current == 'f' {
            if remaining_chars < 4 {
                continue;
            }
            if vec[i + 1] == 'o' && vec[i + 2] == 'u' && vec[i + 3] == 'r' {
                digits.push('4');
                continue;
            }
        }

        if current == 'f' {
            if remaining_chars < 4 {
                continue;
            }
            if vec[i + 1] == 'i' && vec[i + 2] == 'v' && vec[i + 3] == 'e' {
                digits.push('5');
                continue;
            }
        }

        if current == 's' {
            if remaining_chars < 3 {
                continue;
            }
            if vec[i + 1] == 'i' && vec[i + 2] == 'x' {
                digits.push('6');
                continue;
            }
        }

        if current == 's' {
            if remaining_chars < 5 {
                continue;
            }
            if vec[i + 1] == 'e' && vec[i + 2] == 'v' && vec[i + 3] == 'e' && vec[i + 4] == 'n' {
                digits.push('7');
                continue;
            }
        }

        if current == 'e' {
            if remaining_chars < 5 {
                continue;
            }
            if vec[i + 1] == 'i' && vec[i + 2] == 'g' && vec[i + 3] == 'h' && vec[i + 4] == 't' {
                digits.push('8');
                continue;
            }
        }

        if current == 'n' {
            if remaining_chars < 4 {
                continue;
            }
            if vec[i + 1] == 'i' && vec[i + 2] == 'n' && vec[i + 3] == 'e' {
                digits.push('9');
                continue;
            }
        }
    }
    digits
}


