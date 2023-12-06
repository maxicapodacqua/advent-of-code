pub fn process(input: &str) -> String {
    let parsed_input = input
        .lines()
        .collect::<Vec<_>>()
        .iter()
        .map(|line| {
            line.split_whitespace()
                .skip(1)
                .map(|val| val.parse::<i32>().unwrap())
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>();

    let times = &parsed_input[0];
    let distances = &parsed_input[1];

    let mut out = 1;
    for i in 0..times.len() {
        let a = 1.0;
        let b = times[i];
        let c = distances[i];

        let quadratic_base = (b * b) - (4 * c);

        let quadratic_res_top = (b as f64 + (quadratic_base as f64).sqrt()) / 2.0;
        let quadratic_res_bottom = (b as f64 - (quadratic_base as f64).sqrt()) / 2.0;

        let mut result = (quadratic_res_top.floor() - quadratic_res_bottom.ceil()) as i32;
        if result * result == quadratic_base {
            result -= 1;
        } else {
            result += 1;
        }
        dbg!(&result);
        out *= result;
        // dbg!(quadratic_res_top.floor() - quadratic_res_bottom.ceil() + 1.0);

        // result *= quadratic_res_top.floor() - quadratic_res_bottom.ceil() + 1.0;
    }

    out.to_string()
}

mod tests {
    use crate::part1::process;

    #[test]
    fn it_works() {
        let input = r#"Time:      7  15   30
Distance:  9  40  200"#;
        let result = process(input);
        assert_eq!(result, "288");
    }
}
