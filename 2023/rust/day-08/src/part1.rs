use std::{collections::HashMap, str::FromStr};

#[derive(Debug)]
struct Network {
    map: HashMap<String, (String, String)>,
    instructions: Vec<char>,
}

#[derive(Debug, PartialEq, Eq)]
struct ParseNetworkErr;
impl FromStr for Network {
    type Err = ParseNetworkErr;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut lines = s.lines();
        let instructions = lines.next().unwrap().chars().collect::<Vec<char>>();
        // dbg!(&instructions);
        let mut map: HashMap<String, (String, String)> = HashMap::new();
        lines.next(); // skip empty line
        for line in lines {
            let parsed_line = line.split(&['=', ',']).collect::<Vec<_>>();

            let key = parsed_line[0].replace(" ", "");
            let left = parsed_line[1].replace(" (", "");
            let right = parsed_line[2].replace(")", "");
            let right = right.replace(" ", "");

            map.insert(key, (left, right));
        }

        Ok(Network { map, instructions })
    }
}

pub fn process(input: &str) -> String {
    // let input_lines = input.lines();
    // let instructions = input.lines().next().unwrap().chars();

    let network = input.parse::<Network>().unwrap();
    // dbg!(&network);

    let mut steps: usize = 0;

    let mut current_node = network.map.get_key_value("AAA").unwrap();

    for instruction in network.instructions.iter().cycle() {
        // println!(
        //     "Processing instruction {:?} in node {:?}",
        //     &instruction, &current_node
        // );
        if current_node.0 == "ZZZ" {
            return steps.to_string();
        }
        steps += 1;

        current_node = if instruction == &'L' {
            // Moves to the left node
            network.map.get_key_value(&current_node.1 .0).unwrap()
        } else {
            // Moves to the right node
            network.map.get_key_value(&current_node.1 .1).unwrap()
        }
    }

    "".to_string()
}

mod tests {
    use crate::part1::process;

    #[test]
    fn it_works() {
        let input = r#"LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)"#;
        let result = process(input);
        assert_eq!(result, "6");
    }
}
