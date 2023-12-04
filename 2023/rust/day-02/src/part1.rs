pub fn process(input: &str) -> String {
    let limit_red = 12;
    let limit_green = 13;
    let limit_blue = 14;

    let mut possible_games: Vec<i32> = vec![];

    for line in input.lines() {
        let g: Vec<_> = line.split(":").collect();
        let game_id = g
            .first()
            .unwrap()
            .trim()
            .replace("Game ", "")
            .parse::<i32>()
            .unwrap();

        let mut max_red = 0;
        let mut max_blue = 0;
        let mut max_green = 0;

        let subset_cubes = g.last().unwrap().split(";").map(|l| l.trim());

        for subset in subset_cubes {
            subset
                .split(", ")
                .map(|l| {
                    let mut line = l.split(" ");
                    (
                        line.next().unwrap().parse::<i32>().unwrap(),
                        line.last().unwrap(),
                    )
                })
                .for_each(|(val, color)| match color {
                    "red" => {
                        if val > max_red {
                            max_red = val
                        }
                    }
                    "blue" => {
                        if val > max_blue {
                            max_blue = val
                        }
                    }
                    "green" => {
                        if val > max_green {
                            max_green = val
                        }
                    }
                    _ => unreachable!(),
                });
        }
        if max_red <= limit_red && max_blue <= limit_blue && max_green <= limit_green {
            possible_games.push(game_id);
        }
    }

    println!("Possible games {:#?}", possible_games);

    possible_games.iter().fold(0, |acc, x| acc + x).to_string()
}

#[cfg(test)]
mod tests {
    use crate::part1::process;

    #[test]
    fn it_works() {
        let result = process(
            "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
            Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
            Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
            Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
            Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
        );
        assert_eq!(result, "8");
    }
}
