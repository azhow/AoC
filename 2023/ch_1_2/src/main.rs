use std::fs;

fn main() {
    let file_path = "./input";

    let f_contents = fs::read_to_string(file_path)
        .expect("idk file?");

    let nums = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

    let mut acc: u32 = 0;
    let line_split = f_contents.lines();
    for line in line_split {
        let mut max_pos = 0;
        let mut min_pos = std::usize::MAX;
        let mut max_el = std::usize::MAX;
        let mut min_el = std::usize::MAX;
        for (idx, num) in nums.iter().enumerate() {
            let first_occ = line.find(num);
            let last_occ = line.rfind(num);

            if first_occ == None {
                continue;
            }

            if first_occ == last_occ {
                let pos = first_occ.unwrap();

                if pos < min_pos {
                    min_pos = pos;
                    min_el = idx + 1;
                }

                if pos > max_pos {
                    max_pos = pos;
                    max_el = idx + 1;
                }
            }
            else {
                let first_pos = first_occ.unwrap();
                let last_pos = last_occ.unwrap();

                if first_pos < min_pos {
                    min_pos = first_pos;
                    min_el = idx + 1;
                }

                if last_pos > max_pos {
                    max_pos = last_pos;
                    max_el = idx + 1;
                }
            }
        }

        let characters = line.chars();
        let chars_count = characters.clone().count();
        // Find first digit
        for (idx, curr_char) in characters.clone().enumerate() {
            if curr_char.is_numeric() && idx < min_pos {
                min_el = usize::try_from(curr_char.to_digit(10).unwrap()).unwrap();
                break;
            }
        }

        // Find second digit
        for (idx, curr_char) in characters.rev().enumerate() {
            if curr_char.is_numeric() && (chars_count - idx) > max_pos {
                max_el = usize::try_from(curr_char.to_digit(10).unwrap()).unwrap();
                break;
            }
        }

        println!("Line: {line}, min: {min_el}, max: {max_el}");

        let number = min_el.to_string() + &max_el.to_string();

        acc += number.parse::<u32>().unwrap();
    }

    println!("ACC: {acc}");
}
