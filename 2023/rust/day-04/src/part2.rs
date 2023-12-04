pub fn process(input: &str) -> String {
    let grid = input.lines().collect::<Vec<_>>();
    let n_games = grid.len();

    let mut game_history: Vec<i32> = vec![0_i32; n_games];
    for (game, line) in grid.iter().enumerate() {
        let mut parsed_line = line
            .split(":")
            .last()
            .unwrap()
            .split("|")
            .collect::<Vec<_>>();

        let mut winning_numbers = parsed_line
            .first_mut()
            .unwrap()
            .split_whitespace()
            .collect::<Vec<_>>();
        winning_numbers.sort();

        let my_numbers = parsed_line.last().unwrap().split_whitespace();

        let mut count = 0;
        for my_number in my_numbers {
            if winning_numbers.binary_search(&my_number).is_ok() {
                count += 1;
            }
        }
        game_history[game] = count;
    }

    let mut card_counts = vec![1u32; n_games];

    for (game, matches) in game_history.iter().enumerate() {
        let p = *matches;
        let n = card_counts[game];
        card_counts[(game + 1).min(n_games)..(game + p as usize + 1).min(n_games)]
            .iter_mut()
            .for_each(|c| *c += n);
    }
    let res = card_counts.iter().fold(0, |acc, x| acc + x);

    res.to_string()
}

mod tests {
    use crate::part2::process;

    #[test]
    fn it_works() {
        let input = r#"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"#;
        let result = process(input);
        assert_eq!(result, "30");
    }
}
