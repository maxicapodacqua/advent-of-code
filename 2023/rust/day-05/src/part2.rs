use std::{cmp, str::FromStr};

#[derive(Debug)]
struct PlantSource {
    // destination range, source range, range len
    modifiers: Vec<(usize, usize, usize)>,
}
impl PlantSource {
    fn get_source(&self, destination: usize) -> usize {
        for mods in self.modifiers.iter() {
            let (dest_range, source_range, range_len) = mods;
            let range = source_range..&(source_range + range_len);

            if range.contains(&&destination) {
                return (destination - source_range) + dest_range;
            }
        }
        destination
    }
}
#[derive(Debug, PartialEq, Eq)]
struct ParsePlantSourceError;

impl FromStr for PlantSource {
    type Err = ParsePlantSourceError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        // Skip text from top
        let lines = s.lines().skip(1).collect::<Vec<_>>();

        // let mut map = HashMap::new();
        let mut modifiers: Vec<_> = Vec::new();

        for l in lines.iter() {
            let row = l
                .split_whitespace()
                .map(|v| {
                    v.parse::<usize>()
                        .map_err(|_| ParsePlantSourceError)
                        .unwrap()
                })
                .collect::<Vec<_>>();

            let destination_range = row[0];
            let source_range = row[1];
            let range_length = row[2];
            modifiers.push((destination_range, source_range, range_length));
        }

        Ok(PlantSource { modifiers })
    }
}
// Parsed seed line by creating tuples that represent (start, end) of an
// inclusive Range start..=range.
// Sorts the tuples by start value and then proceeds to merge intervals to avoid
// parsing the same seed twice
fn parse_seed_line(input: &str) -> Vec<(usize, usize)> {
    let seed_iter = input
        .split_whitespace()
        .skip(1)
        .map(|val| val.parse::<usize>().unwrap())
        .collect::<Vec<_>>();
    let seed_by_2 = seed_iter.chunks(2);
    // dbg!(&seed_by_2.collect::<Vec<_>>());
    let mut sorted_seeds: Vec<(usize, usize)> = vec![];
    for seed_data in seed_by_2 {
        sorted_seeds.push((seed_data[0], (seed_data[0] + seed_data[1] - 1)));
    }
    sorted_seeds.sort_by(|a, b| a.0.cmp(&b.0));

    let mut joined_seeds: Vec<(usize, usize)> = vec![];
    joined_seeds.push(sorted_seeds[0].clone());
    for i in 1..sorted_seeds.len() {
        let curr = sorted_seeds[i].clone();
        let j = joined_seeds.len() - 1;

        if curr.0 >= joined_seeds[j].0 && curr.0 <= joined_seeds[j].1 {
            joined_seeds[j].1 = cmp::max(curr.1, joined_seeds[j].1);
        } else {
            joined_seeds.push(curr);
        }
    }
    joined_seeds
}

pub fn process(input: &str) -> String {
    let seeds_line = input.lines().next().unwrap();
    let parsed_seed_line = parse_seed_line(&seeds_line);

    let mut body = input
        .split("\n\n")
        .skip(1) // skip seeds line
        .map(|block| block.parse::<PlantSource>().unwrap());
    let soil = body.next().unwrap();
    let fertilizer = body.next().unwrap();
    let water = body.next().unwrap();
    let light = body.next().unwrap();
    let temperature = body.next().unwrap();
    let humidity = body.next().unwrap();
    let location = body.next().unwrap();

    let mut min_location = usize::MAX;

    for (start, end) in parsed_seed_line.iter() {
        for seed in *start..=*end {
            let s = soil.get_source(seed);
            let f = fertilizer.get_source(s);
            let w = water.get_source(f);
            let l = light.get_source(w);
            let t = temperature.get_source(l);
            let h = humidity.get_source(t);
            let loc = location.get_source(h);

            min_location = min_location.min(loc);
        }
    }

    min_location.to_string()
}

mod tests {

    use crate::part2::{process, PlantSource};

    use super::parse_seed_line;

    #[test]
    fn parse_seed_line_works() {
        let input = r#"seeds: 79 14 55 13"#;
        let res = parse_seed_line(input);
        assert_eq!(res.len(), 2);
        assert_eq!(res[0], (55, 67));
        assert_eq!(res[1], (79, 92));

        let input = r#"seeds: 79 14 55 33"#;
        let res = parse_seed_line(input);
        assert_eq!(res.len(), 1);
        assert_eq!(res[0], (55, 92));

        // dbg!(res);
    }
    #[test]
    fn plant_from_str_works() {
        let input = r#"seed-to-soil map:
50 98 2
52 50 48"#;
        let soil = input.parse::<PlantSource>().unwrap();
        assert_eq!(soil.get_source(98), 50);
        assert_eq!(soil.get_source(53), 55);
        assert_eq!(soil.get_source(2), 2);
    }

    #[test]
    fn plant_get_source_works() {
        let mut mods = Vec::new();
        mods.push((50, 98, 2));
        mods.push((52, 50, 48));
        let soil = PlantSource { modifiers: mods };
        assert_eq!(soil.get_source(53), 55);
        assert_eq!(soil.get_source(2), 2);
    }
    #[test]
    fn it_works() {
        let input = r#"seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"#;
        let result = process(input);
        assert_eq!(result, "46");
    }
}
