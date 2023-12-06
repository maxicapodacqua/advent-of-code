pub fn process(input: &str) -> String {
    let parsed_input = input
        .lines()
        .collect::<Vec<_>>()
        .iter()
        .map(|line| {
            line.replace("Time:", "")
                .replace("Distance:", "")
                .chars()
                .filter(|c| !c.is_whitespace())
                .collect::<String>()
        })
        .collect::<Vec<_>>();
    let b = &parsed_input[0].parse::<f64>().unwrap();
    let c = &parsed_input[1].parse::<f64>().unwrap();

    let quadratic_base = b.powf(2.0) - (4.0 * c);

    let quadratic_res_top = (*b as f64 + (quadratic_base as f64).sqrt()) / 2.0;
    let quadratic_res_bottom = (*b as f64 - (quadratic_base as f64).sqrt()) / 2.0;

    let mut result = quadratic_res_top.floor() - quadratic_res_bottom.ceil();
    if result * result == quadratic_base {
        result -= 1.0;
    } else {
        result += 1.0;
    }

    result.to_string()
}

mod tests {
    use crate::part2::process;

    #[test]
    fn it_works() {
        let input = r#"Time:      7  15   30
Distance:  9  40  200"#;
        let result = process(input);
        assert_eq!(result, "71503");
    }
}
