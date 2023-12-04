use std::collections::HashMap;

pub fn process(input: &str) -> String {
    // Iterate over grid
    // Skip . and numeric
    // Check surroundings for numbers and add the first digit coordinates in a vector (if number is 123, store the coordinates of 1)
    // Iterate in coordinates found
    // Capture the rest of the digits (from coordinates of 1, concat 2 and 3, to create 123 back)
    // Add digits in accum
    // Return accum
    let grid = input.lines().collect::<Vec<&str>>();

    let mut coords: HashMap<(usize, usize), bool> = HashMap::new();

    // let mut coordinates: Vec<(usize, usize)> = Vec::new();

    for (row, line) in grid.iter().enumerate() {
        for (col, char) in line.chars().enumerate() {
            if char.is_digit(10) || char == '.' {
                continue;
            }

            // Check surroundings
            for i in row - 1..=row + 1 {
                for j in col - 1..=col + 1 {
                    // Skip out of bounds
                    if i > grid.len() || j > line.len() {
                        continue;
                    }
                    // Skip non numeric in case it's adjacent to other symbol
                    if !grid[i].chars().nth(j).unwrap().is_digit(10) {
                        continue;
                    }

                    // Find the first digit of number
                    let mut first_digit_coordinate = j;
                    while first_digit_coordinate > 0
                        && grid[i]
                            .chars()
                            .nth(first_digit_coordinate - 1)
                            .unwrap()
                            .is_digit(10)
                    {
                        first_digit_coordinate -= 1;
                    }

                    // let a = coord.get(&(i, first_digit_coordinate)).unwrap();
                    // let exists = coord.get(&(i, first_digit_coordinate)).unwrap_or(&false);

                    // let exists = coordinates
                    //     .iter()
                    //     .any(|(x, y)| x == &i && y == &first_digit_coordinate);
                    // if !exists {
                    coords.insert((i, first_digit_coordinate), true);
                    // coordinates.push((i, first_digit_coordinate));
                    // }
                }
            }
        }
    }

    let mut acc = 0;

    for ((x, y), _) in coords.iter() {
        let mut digits: Vec<String> = Vec::new();

        // Find next digits
        let mut temp_y = *y;
        while temp_y < grid[*x].len() && grid[*x].chars().nth(temp_y).unwrap().is_digit(10) {
            digits.push(grid[*x].chars().nth(temp_y).unwrap().to_string());
            temp_y += 1;
        }
        acc += digits.join("").parse::<usize>().unwrap();
    }

    acc.to_string()
}

mod tests {
    use crate::part1::process;

    #[test]
    fn it_works() {
        let input = r#"467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."#;
        let result = process(input);
        assert_eq!(result, "4361");
    }
}
