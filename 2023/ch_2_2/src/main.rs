use std::fs;

#[derive(Debug)]
struct Game {
    power: usize,
}

impl Game {
    fn new(input: String) -> Game {
        let mut col_split = input.split(':');

        let token = col_split.next().unwrap();
        let _id = token.split(' ').last().unwrap().parse::<usize>().unwrap();

        let mut min_red = 0;
        let mut min_green = 0;
        let mut min_blue = 0;

        let semi_col_split = col_split.next().unwrap().split(';');
        for cubes in semi_col_split {
            let cubes_split = cubes.split(',');
            for cube_num_split in cubes_split {
                let mut space_split = cube_num_split.split(' ');
                space_split.next();
                let number = space_split.next().unwrap().parse::<usize>().unwrap();
                let color = space_split.next().unwrap();

                match color {
                    "red" => min_red = if number > min_red { number } else { min_red },
                    "green" => min_green = if number > min_green { number } else { min_green },
                    "blue" => min_blue = if number > min_blue { number } else { min_blue },
                    _ => (),
                }
            }
        }

        Game { power: min_red * min_green * min_blue }
    }
}

fn main() {
    let file_path = "./input";

    let f_contents = fs::read_to_string(file_path)
        .expect("idk file?");

    let mut acc = 0;
    let line_split = f_contents.lines();
    for line in line_split {
        let curr_game = Game::new(line.to_string());
        println!("Valid Game {curr_game:?}");
        acc += curr_game.power;
    }

    println!("Sum of valid game ids: {acc}");
}
