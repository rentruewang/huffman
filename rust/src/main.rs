use std::cmp::{Eq, Ord, PartialEq, PartialOrd};
use std::collections::{BinaryHeap, HashMap};
use std::{env, fs};

#[derive(Eq, PartialEq, PartialOrd, Ord)]
struct Node<K, V> {
    val: V,
    key: K,
    left: Option<Box<Node<K, V>>>,
    right: Option<Box<Node<K, V>>>,
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let content: Vec<char> = fs::read_to_string(&args[1]).unwrap().chars().collect();

    let mut word_count: HashMap<char, isize> = HashMap::new();

    for ch in content.into_iter() {
        let v = word_count.get_mut(&ch);
        if let Some(v) = v {
            *v += 1;
        } else {
            word_count.insert(ch, 1);
        }
    }

    let mut pq: BinaryHeap<Node<char, isize>> = BinaryHeap::new();

    for (key, val) in word_count.into_iter() {
        let s = Node {
            key,
            val: -val,
            left: None,
            right: None,
        };

        pq.push(s);
    }

    let head: Node<char, isize>;

    loop {
        let (f, s) = (pq.pop(), pq.pop());
        let first = f.unwrap();
        let second = if let Some(s) = s {
            s
        } else {
            head = first;
            break;
        };

        let nn = Node {
            key: '\0',
            val: first.val + second.val,
            left: Some(Box::new(first)),
            right: Some(Box::new(second)),
        };

        pq.push(nn);
    }

    let result = traverse(head);

    println!("{:?}", result);
}

// move hash map out of this function after use
fn traverse(node: Node<char, isize>) -> HashMap<char, String> {
    fn _traverse(node: Node<char, isize>, map: &mut HashMap<char, String>, path: &String) {
        let left = node.left;
        let right = node.right;

        let pc = path.clone();
        if let Some(left) = left {
            _traverse(*left, map, &(pc + "0"));
        } else {
            map.insert(node.key, pc);
        }

        let pc = path.clone();
        if let Some(right) = right {
            _traverse(*right, map, &(pc + "1"));
        } else {
            map.insert(node.key, pc);
        }
    }
    let mut path_list: HashMap<char, String> = HashMap::new();
    _traverse(node, &mut path_list, &String::from(""));
    path_list
}
