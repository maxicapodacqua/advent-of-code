use std::{collections::HashMap, str::FromStr};

#[derive(Debug, Ord, PartialOrd, Eq, PartialEq)]
enum HandType {
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}
#[derive(Debug, PartialOrd, Eq, PartialEq)]
struct Hand {
    cards: String,
    bid: usize,
    hand_type: HandType,
}

impl Hand {
    fn process_hand_type(cards: &String) -> HandType {
        let chars = &cards.chars().collect::<Vec<_>>();
        let has_joker = chars.contains(&'J');

        let chars_map = chars.iter().fold(HashMap::new(), |mut acc, val| {
            let entry = acc.entry(val).or_insert(0);
            *entry += 1;
            acc
        });
        let mut sorted_chars = chars_map.iter().collect::<Vec<_>>();

        if has_joker {
            let max_char = chars_map
                .iter()
                .filter(|c| **c.0 != 'J')
                .max_by(|a, b| a.1.cmp(&b.1));
            if let Some(max_char) = max_char {
                let replaced = &cards.replacen("J", &max_char.0.to_string(), 1);
                // dbg!(replaced);
                return Hand::process_hand_type(replaced);
            } else {
                // continue processing if there is no max char
            }
        }

        sorted_chars.sort_by(|a, b| b.1.cmp(&a.1));

        let hand_type: HandType = match sorted_chars[0].1 {
            5 => HandType::FiveOfAKind,
            4 => HandType::FourOfAKind,
            3 => {
                if *sorted_chars[1].1 == 2 {
                    HandType::FullHouse
                } else {
                    HandType::ThreeOfAKind
                }
            }
            2 => {
                if *sorted_chars[1].1 == 2 {
                    HandType::TwoPair
                } else {
                    HandType::OnePair
                }
            }
            1 => HandType::HighCard,
            _ => unreachable!(),
        };

        hand_type
    }
}

#[derive(Debug, PartialEq, Eq)]
struct ParseHandErr;
impl FromStr for Hand {
    type Err = ParseHandErr;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut iter = s.split_whitespace();

        let cards = iter.next().unwrap().to_string();
        let bid = iter
            .last()
            .unwrap()
            .parse::<usize>()
            .map_err(|_| ParseHandErr)
            .unwrap();

        let hand_type = Hand::process_hand_type(&cards);
        Ok(Hand {
            cards,
            bid,
            hand_type,
        })
    }
}

impl Ord for Hand {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        let iter = self.cards.chars().zip(other.cards.chars()).into_iter();

        let value_map = HashMap::from([
            ('A', 13),
            ('K', 12),
            ('Q', 11),
            ('T', 10),
            ('9', 9),
            ('8', 8),
            ('7', 7),
            ('6', 6),
            ('5', 5),
            ('4', 4),
            ('3', 3),
            ('2', 2),
            ('J', 1), // now J is lower
        ]);
        for (a, b) in iter {
            let a_val = value_map.get(&a).unwrap();
            let b_val = value_map.get(&b).unwrap();
            let cmp_res = a_val.cmp(&b_val);
            if !cmp_res.is_eq() {
                return cmp_res;
            }
        }
        unreachable!("Compared none stop and both hands are equals");
    }
}

pub fn process(input: &str) -> String {
    let mut hands: Vec<Hand> = input.lines().map(|l| l.parse::<Hand>().unwrap()).collect();

    // dbg!(&hands);
    hands.sort_by(|a, b| {
        let cmp_res = a.hand_type.cmp(&b.hand_type);
        if cmp_res.is_eq() {
            a.cmp(&b)
        } else {
            cmp_res
        }
    });

    let res = hands
        .iter()
        .enumerate()
        .fold(0, |acc, (rank, hand)| acc + (rank + 1) * hand.bid);

    // dbg!(&hands);
    res.to_string()
}

mod tests {
    use crate::part2::process;

    #[test]
    fn it_works() {
        let input = r#"32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483"#;
        let result = process(input);
        assert_eq!(result, "5905");
    }
}
