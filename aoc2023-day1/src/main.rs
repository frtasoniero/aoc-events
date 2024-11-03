use std::fs::File;
use std::io::{self, BufRead};

fn convert_word_to_number(line: String) -> String {
    let word_to_digit = [
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("eight", "8"),
        ("nine", "9"),
    ];

    let mut result = String::new();
    let mut i = 0;

    while i < line.len() {
        let mut matched = false;

        for (word, digit) in &word_to_digit {
            if line[i..].starts_with(word) {
                result.push_str(digit);
                i += word.len() - 1;
                matched = true;
                break;
            }
        }

        if !matched {
            result.push(line[i..].chars().next().unwrap());
            i += 1;
        }
    }

    result
}

fn check_number_inside_string(line: String) -> i32 {
    let mut number_as_string = String::new();

    let converted_line = convert_word_to_number(line);

    for character in converted_line.chars() {
        if character.is_digit(10) {
            number_as_string.push(character);
            break;
        }
    }

    let reversed_str: String = converted_line.chars().rev().collect();

    for character in reversed_str.chars() {
        if character.is_digit(10) {
            number_as_string.push(character);
            break;
        }
    }

    let number_as_int: i32 = match number_as_string.parse::<i32>() {
        Ok(num) => num,
        Err(_) => 0,
    };

    number_as_int
}

fn main() -> io::Result<()> {
    let file = File::open("./src/input.txt")?;

    let reader = io::BufReader::new(file);

    let mut result: i32 = 0;

    for line in reader.lines() {
        let line = line?;
        // let line = "xtwone3four".to_string();
        println!("{line}");
        let value: i32 = check_number_inside_string(line);
        println!("\t{value}\n");
        result = result + value;
    }

    println!("The result is {result}");

    Ok(())
}
