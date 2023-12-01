fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
    let mut acc = 0;

    for line in input.lines() {
        let numerics = line
            .match_indices(char::is_numeric)
            .map(|(i, v)| (i, v.parse::<i32>().unwrap()))
            .into_iter();
        let ones = line.match_indices("one").map(|(i, _)| (i, 1)).into_iter();
        let two = line.match_indices("two").map(|(i, _)| (i, 2)).into_iter();
        let three = line.match_indices("three").map(|(i, _)| (i, 3)).into_iter();
        let four = line.match_indices("four").map(|(i, _)| (i, 4)).into_iter();
        let five = line.match_indices("five").map(|(i, _)| (i, 5)).into_iter();
        let six = line.match_indices("six").map(|(i, _)| (i, 6)).into_iter();
        let seven = line.match_indices("seven").map(|(i, _)| (i, 7)).into_iter();
        let eight = line.match_indices("eight").map(|(i, _)| (i, 8)).into_iter();
        let nine = line.match_indices("nine").map(|(i, _)| (i, 9)).into_iter();

        let mut chained: Vec<_> = numerics
            .chain(ones)
            .chain(two)
            .chain(three)
            .chain(four)
            .chain(five)
            .chain(six)
            .chain(seven)
            .chain(eight)
            .chain(nine)
            .into_iter()
            .collect();
        chained.sort();

        let first = chained.first().unwrap();

        let last = chained.last().unwrap_or(first);
        acc += (first.1 * 10) + last.1;

        // dbg!(first, last, acc);
    }

    acc.to_string()
}

mod tests {
    use crate::part2;

    #[test]
    fn it_works() {
        let result = part2(
            "two1nine
            eightwothree
            abcone2threexyz
            xtwone3four
            4nineeightseven2
            zoneight234
            7pqrstsixteen",
        );
        assert_eq!(result, "281");
    }
}
