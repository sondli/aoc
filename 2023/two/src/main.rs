use std::fs;

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Something went wrong reading the file");
    q1(&contents);
    q2(&contents);
}

fn q1(content: &String) {
    const MAX_RED: i32 = 12;
    const MAX_GREEN: i32 = 13;
    const MAX_BLUE: i32 = 14;

    let mut possible_games: Vec<i32> = Vec::new();

    for line in content.lines() {
        let line_split: Vec<&str> = line.split(":").collect();
        let game_number: i32 = line_split[0].split(" ").collect::<Vec<&str>>()[1]
            .parse()
            .unwrap();

        let sets: Vec<&str> = line_split[1].split(";").collect();

        let game_impossible = sets.iter().any(|set| {
            let set = parse_set(set);
            set.red > MAX_RED || set.green > MAX_GREEN || set.blue > MAX_BLUE
        });

        if game_impossible {
            continue;
        }

        possible_games.push(game_number);
    }
    let sum: i32 = possible_games.iter().sum();
    println!("The sum of the possible games is {}", sum);
}

fn q2(content: &String) {
    let mut powers: Vec<i32> = Vec::new();
    for line in content.lines() {
        let line_split: Vec<&str> = line.split(":").collect();
        let sets: Vec<Set> = line_split[1]
            .split(";")
            .collect::<Vec<&str>>()
            .iter()
            .map(|set| parse_set(set))
            .collect();

        let mut smallest_red = sets.iter().map(|set| set.red).max().unwrap();
        let mut smallest_green = sets.iter().map(|set| set.green).max().unwrap();
        let mut smallest_blue = sets.iter().map(|set| set.blue).max().unwrap();

        if smallest_red == 0 {
            smallest_red = 1;
        }

        if smallest_green == 0 {
            smallest_green = 1;
        }

        if smallest_blue == 0 {
            smallest_blue = 1;
        }

        let power = smallest_red * smallest_green * smallest_blue;
        powers.push(power);
    }

    let sum = powers.iter().sum::<i32>();
    println!("The sum of the powers is {}", sum);
}

struct Set {
    red: i32,
    green: i32,
    blue: i32,
}

fn parse_set(set: &str) -> Set {
    let set_split: Vec<&str> = set.split(",").collect();
    let mut red = 0;
    let mut green = 0;
    let mut blue = 0;
    for i in 0..set_split.len() {
        let cube_info: Vec<&str> = set_split[i]
            .split(" ")
            .filter(|x| x.to_owned() != "")
            .collect();
        let cube_number: i32 = cube_info[0].parse().unwrap();
        let cube_color = cube_info[1].to_string();
        match cube_color.as_str() {
            "red" => red = cube_number,
            "green" => green = cube_number,
            "blue" => blue = cube_number,
            _ => println!("Invalid color"),
        }
    }

    Set { red, green, blue }
}
