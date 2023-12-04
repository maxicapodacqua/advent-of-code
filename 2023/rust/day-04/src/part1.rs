pub fn process(input: &str) -> String {
    let grid = input.lines().collect::<Vec<_>>();

    let mut acc = 0;

    // Iterate per line
    // extract left side of pipeline (winning numbers) and right side of pipeline (my numbers)
    // sort winning numbers
    // iterate on my numbers
    //    binary search my numbers in winning numbers
    //    if found, add result to vec
    // calculate total based on vec size

    for line in grid {
        let mut matched_numbers: Vec<_> = Vec::new();
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

        for my_number in my_numbers {
            if winning_numbers.binary_search(&my_number).is_ok() {
                matched_numbers.push(my_number);
            }
        }
        if !matched_numbers.is_empty() {
            let res_line = 2_usize.pow((matched_numbers.len() - 1).try_into().unwrap());
            acc += res_line;
        }
    }

    acc.to_string()
}

mod tests {
    use crate::part1::process;

    #[test]
    fn it_works() {
        let input = r#"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"#;
        let result = process(input);
        assert_eq!(result, "13");
    }
}
