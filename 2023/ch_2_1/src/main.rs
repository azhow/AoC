use std::fs;

#[derive(Debug)]
struct Game {
    id: usize,
    red: usize,
    green: usize,
    blue: usize,
    valid: bool,
}

impl Game {
    fn new(input: String) -> Game {
        let max_red = 12;
        let max_green = 13;
        let max_blue = 14;

        let mut col_split = input.split(':');

        let token = col_split.next().unwrap();
        let id = token.split(' ').last().unwrap().parse::<usize>().unwrap();

        let red = 0;
        let green = 0;
        let blue = 0;
        let mut valid = true;

        let semi_col_split = col_split.next().unwrap().split(';');
        for cubes in semi_col_split {
            let cubes_split = cubes.split(',');
            for cube_num_split in cubes_split {
                let mut space_split = cube_num_split.split(' ');
                space_split.next();
                let number = space_split.next().unwrap().parse::<usize>().unwrap();
                let color = space_split.next().unwrap();

                match color {
                    "red" => valid = valid && number <= max_red,
                    "green" => valid = valid && number <= max_green,
                    "blue" => valid = valid && number <= max_blue,
                    _ => (),
                }
            }
        }

        Game { id, red, green, blue, valid }
    }
}

fn main() {
    let file_path = "./input";

    let f_contents = fs::read_to_string(file_path)
        .expect("idk file?");

    let mut valid_games = Vec::new();
    let line_split = f_contents.lines();
    for line in line_split {
        let curr_game = Game::new(line.to_string());
        if curr_game.valid {
            valid_games.push(curr_game.id);
            println!("Valid Game {curr_game:?}");
        }
    }

    let mut acc = 0;
    valid_games.iter().for_each(|g| acc += g );
    println!("Sum of valid game ids: {acc}");
}
