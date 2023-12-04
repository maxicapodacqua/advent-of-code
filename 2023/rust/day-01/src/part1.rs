use wasm_bindgen::prelude::*;

#[wasm_bindgen(js_name = processPart1)]
pub fn process(input: &str) -> String {
    let lines = input.lines();
    let mut acc = 0;
    let mut first_digit: i32 = -1;
    let mut last_digit = -1;
    for l in lines {
        let chars = l.chars();
        for c in chars {
            if c.is_digit(10) {
                if first_digit == -1 {
                    first_digit = c.to_digit(10).unwrap() as i32;
                } else {
                    last_digit = c.to_digit(10).unwrap() as i32;
                }
            }
        }
        // When there is only one digit in the line
        if last_digit == -1 {
            last_digit = first_digit
        }
        acc += (first_digit * 10) + last_digit;
        // dbg!(l, first_digit, last_digit, acc);
        last_digit = -1;
        first_digit = -1;
    }

    acc.to_string()
}

#[cfg(test)]
mod tests {
    use crate::part1::process;

    #[test]
    fn it_works() {
        let result = process(
            "1abc2
        pqr3stu8vwx
        a1b2c3d4e5f
        treb7uchet",
        );
        assert_eq!(result, "142");
    }
}
