use std::{cmp, collections, env, fs};

#[derive(cmp::Eq, cmp::PartialEq, cmp::PartialOrd, cmp::Ord)]
struct Node<K, V> {
    val: V,
    key: K,
    left: Option<Box<Node<K, V>>>,
    right: Option<Box<Node<K, V>>>,
}

// struct Node<K, V> {
//     key: K,
//     val: V,
//     left: Option<Box<Node<K, V>>>,
//     right: Option<Box<Node<K, V>>>,
// }

// impl<K, V> cmp::Eq for Node<K, V> where V: cmp::Ord {}
// impl<K, V> cmp::PartialEq for Node<K, V>
// where
//     V: cmp::Ord,
// {
//     fn eq(&self, other: &Self) -> bool {
//         self.val == other.val
//     }
//     fn ne(&self, other: &Self) -> bool {
//         self.val != other.val
//     }
// }

// impl<K, V> cmp::Ord for Node<K, V>
// where
//     V: cmp::Ord,
// {
//     fn cmp(&self, other: &Self) -> cmp::Ordering {
//         if let Some(a) = self.partial_cmp(other) {
//             a
//         } else {
//             unreachable!()
//         }
//     }
//     fn max(self, other: Self) -> Self {
//         cmp::max(self, other)
//     }
//     fn min(self, other: Self) -> Self {
//         cmp::min(self, other)
//     }
//     // fn clamp(self, min: Self, max: Self) -> Self {
//     //     cmp::min(max, cmp::max(min, self))
//     // }
// }

// impl<K, V> cmp::PartialOrd for Node<K, V>
// where
//     V: cmp::Ord,
// {
//     fn partial_cmp(&self, other: &Self) -> Option<cmp::Ordering> {
//         if self.lt(other) {
//             Some(cmp::Ordering::Less)
//         } else if self.ge(other) {
//             Some(cmp::Ordering::Greater)
//         } else {
//             Some(cmp::Ordering::Equal)
//         }
//     }
//     fn lt(&self, other: &Self) -> bool {
//         self.val < other.val
//     }
//     fn le(&self, other: &Self) -> bool {
//         self.val <= other.val
//     }
//     fn gt(&self, other: &Self) -> bool {
//         self.val > other.val
//     }
//     fn ge(&self, other: &Self) -> bool {
//         self.val >= other.val
//     }
// }

fn main() {
    let args: Vec<String> = env::args().collect();
    let content: Vec<char> = fs::read_to_string(&args[1]).unwrap().chars().collect();

    let mut word_count: collections::HashMap<char, isize> = collections::HashMap::new();

    for ch in content.into_iter() {
        let v = word_count.get_mut(&ch);
        if let Some(v) = v {
            *v += 1;
        } else {
            word_count.insert(ch, 1);
        }
    }

    let mut pq: collections::BinaryHeap<Node<char, isize>> = collections::BinaryHeap::new();

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
fn traverse(node: Node<char, isize>) -> collections::HashMap<char, String> {
    fn _traverse(
        node: Node<char, isize>,
        map: &mut collections::HashMap<char, String>,
        path: &String,
    ) {
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
    let mut path_list: collections::HashMap<char, String> = collections::HashMap::new();
    _traverse(node, &mut path_list, &String::from(""));
    path_list
}
