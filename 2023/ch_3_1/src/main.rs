use std::fs;

fn main() {
    let file_path = "./input";

    let f_contents = fs::read_to_string(file_path)
        .expect("idk file?");

    let mut parts_positions = Vec::new();
    let mut parts = Vec::new();

    let line_split = f_contents.lines();
    for (line_idx, line) in line_split.enumerate() {
        for (col_idx, curr_char) in line.chars().enumerate() {
            if !curr_char.is_ascii_digit() && curr_char != '.' {
                parts_positions.push((line_idx - 1, col_idx - 1));
                parts_positions.push((line_idx - 1, col_idx));
                parts_positions.push((line_idx - 1, col_idx + 1));
                parts_positions.push((line_idx, col_idx - 1));
                parts_positions.push((line_idx, col_idx + 1));
                parts_positions.push((line_idx + 1, col_idx - 1));
                parts_positions.push((line_idx + 1, col_idx));
                parts_positions.push((line_idx + 1, col_idx + 1));
            }
        }
    }

    let line_split = f_contents.lines();
    for (line_idx, line) in line_split.enumerate() {
        let mut number = String::new();
        let mut valid_number = false;
        for (col_idx, curr_char) in line.chars().enumerate() {
            if curr_char.is_ascii_digit() {
                number.push(curr_char);
                valid_number = parts_positions.contains(&(line_idx, col_idx)) || valid_number;
                if col_idx == line.len() - 1 {
                    if valid_number {
                        parts.push(number.parse::<u32>().unwrap());
                    }
                }
            }
            else {
                if valid_number {
                    parts.push(number.parse::<u32>().unwrap());
                }
                number = "".to_string();
                valid_number = false;
            }
        }
    }

    let acc_parts: u32 = parts.iter().sum();
    println!("Acc: {acc_parts}");
}
