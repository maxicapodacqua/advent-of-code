use std::collections::HashMap;

pub fn process(input: &str) -> String {
    let grid = input.lines().collect::<Vec<&str>>();

    let mut coords: HashMap<(usize, usize), HashMap<(usize, usize), bool>> = HashMap::new();

    // let mut coordinates: Vec<(usize, usize)> = Vec::new();

    for (row, line) in grid.iter().enumerate() {
        for (col, char) in line.chars().enumerate() {
            if char != '*' {
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
                    let mut new_hm: HashMap<(usize, usize), bool> = HashMap::new();
                    new_hm.insert((i, first_digit_coordinate), true);

                    let matches = coords.entry((row, col)).or_insert(new_hm);
                    matches.insert((i, first_digit_coordinate), true);
                    // matches.push((i, first_digit_coordinate));
                    // Increase match counter
                    // *count = (count.0, count.1, count.2 + 1);
                    // coordinates.push((i, first_digit_coordinate));
                    // }
                }
            }
        }
    }

    let mut acc = 0;

    // dbg!(coords);
    for (gear_coords, matches) in coords.iter() {
        if matches.len() != 2 {
            continue;
        }

        let mut gear_ratio = 1;

        for ((x, y), _) in matches.iter() {
            let mut digits: Vec<String> = Vec::new();

            // Find next digits
            let mut temp_y = *y;
            while temp_y < grid[*x].len() && grid[*x].chars().nth(temp_y).unwrap().is_digit(10) {
                digits.push(grid[*x].chars().nth(temp_y).unwrap().to_string());
                temp_y += 1;
            }
            gear_ratio *= digits.join("").parse::<usize>().unwrap();
            // dbg!(digits.join("").parse::<usize>().unwrap());
        }
        acc += gear_ratio;
        gear_ratio = 0;

        // dbg!(gear_coords, matches.len());
    }

    // for ((x, y), _) in coords.iter() {
    //     let mut digits: Vec<String> = Vec::new();

    //     // Find next digits
    //     let mut temp_y = *y;
    //     while temp_y < grid[*x].len() && grid[*x].chars().nth(temp_y).unwrap().is_digit(10) {
    //         digits.push(grid[*x].chars().nth(temp_y).unwrap().to_string());
    //         temp_y += 1;
    //     }
    //     acc += digits.join("").parse::<usize>().unwrap();
    // }

    acc.to_string()
}

mod tests {
    use crate::part2::process;

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
        assert_eq!(result, "467835");
    }
}
