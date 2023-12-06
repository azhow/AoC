use std::fs;

fn main() {
    let file_path = "./input";

    let f_contents = fs::read_to_string(file_path)
        .expect("idk file?");

    let mut acc: u32 = 0;
    let line_split = f_contents.lines();
    for line in line_split {
        let mut number = String::new();
        let characters = line.chars();
        // Find first digit
        for curr_char in characters.clone() {
            if curr_char.is_numeric() {
                number.push(curr_char);
                break;
            }
        }
        // Find second digit
        for curr_char in characters.rev() {
            if curr_char.is_numeric() {
                number.push(curr_char);
                break;
            }
        }

        acc += number.parse::<u32>().unwrap();
    }

    println!("ACC: {acc}");
}
