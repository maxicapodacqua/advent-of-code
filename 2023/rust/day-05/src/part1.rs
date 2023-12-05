use std::{collections::HashMap, str::FromStr};

#[derive(Debug)]
struct PlantSource {
    // destination, source
    // map: HashMap<usize, usize>,
    modifiers: Vec<(usize, usize, usize)>,
}
impl PlantSource {
    fn get_source(&self, destination: usize) -> usize {
        for mods in self.modifiers.iter() {
            let (dest_range, source_range, range_len) = mods;
            let range = source_range..&(source_range + range_len);

            if range.contains(&&destination) {
                // dbg!((destination - source_range) + dest_range);
                return (destination - source_range) + dest_range;
            }
        }
        destination
        // let res = self.map.get(&destination).unwrap_or(&destination);
        // *res
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

            // for i in 0..range_length {
            //     map.insert(source_range + i, destination_range + i);
            // }
        }

        Ok(PlantSource { modifiers })
    }
}

pub fn process(input: &str) -> String {
    let seeds_line = input.lines().next().unwrap();

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

    for seed in seeds_line
        .split_whitespace()
        .skip(1)
        .map(|seed| seed.parse::<usize>().unwrap())
    {
        let s = soil.get_source(seed);
        let f = fertilizer.get_source(s);
        let w = water.get_source(f);
        let l = light.get_source(w);
        let t = temperature.get_source(l);
        let h = humidity.get_source(t);
        let loc = location.get_source(h);

        min_location = min_location.min(loc);
    }

    min_location.to_string()
}

mod tests {
    use std::collections::HashMap;

    use crate::part1::{process, PlantSource};

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
        assert_eq!(result, "35");
    }
}
