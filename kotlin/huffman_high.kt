import java.util.PriorityQueue
import kotlin.collections.HashMap
import kotlin.collections.List

typealias nodechar = Node<Char>

fun main() {
    val string: String = "hello world!"
    val wc = word_count(string)
    val t = tree(wc)
    val code = walk(t)
    println(code.toList().sortedBy{
        val (key, _) = it
        key
    })
    //[( , 1100), (!, 1101), (d, 000), (e, 001), (h, 0110), (l, 10), (o, 111), (r, 0111), (w, 010)]
}

fun word_count(string: String): HashMap<Char, Int> {
    val charlist = string.toList()
    return charlist.fold(
        initial=HashMap<Char, Int>(),
        operation={
            map, elem ->
            if (map.containsKey(elem))
                map[elem] = map[elem]!! + 1
            else
                map[elem] = 1
            map
        }
    )
}

fun tree(wc: HashMap<Char, Int>): nodechar {
    val list: List<Pair<Char, Int>> = wc.toList()
    val node_list = list.map {
        val (item, count) = it
        Node<Char>(null, item, count)
    }.toList()
    val priority_queue = PriorityQueue<nodechar>(node_list.size, object: Comparator<nodechar> {
        override fun compare(self: nodechar, other: nodechar): Int {
            val (_, _, selfcnt) = self
            val (_, _, othercnt) = other
            return selfcnt-othercnt
        }
    })
    node_list.forEach{priority_queue.add(it)}

    while (priority_queue.size >= 1) {
        val first = priority_queue.poll()
        val second = priority_queue.poll() ?: return first
        val node = nodechar(Pair(first, second), '\\', first.count+second.count)
        priority_queue.add(node)
    }
    // this is unreachable
    return nodechar(null, '\\', 0)
}

fun walk(node: nodechar): HashMap<Char, String> {
    val hashmap = HashMap<Char, String>()
    fun walk_rec(node: nodechar, path: String = ""): Unit {
        val (subtree, item, _) = node
        val (left, right) = if (subtree==null) {
            hashmap[item] = path
            return
        } else {
            subtree
        }
        walk_rec(left, path+"0")
        walk_rec(right, path+"1")
    }
    walk_rec(node)
    return hashmap
}

data class Node<E>(val subtree: Pair<Node<E>, Node<E>>?, val item: E, val count: Int)