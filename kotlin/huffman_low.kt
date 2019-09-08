import java.util.*

fun main() {
    val string: String = "hello world!"
    val wc = word_count(string)
    val bt = build_tree(wc)
    val tv = traversal(bt!!)
    println(tv.sortedBy{
        val (key, _) = it
        key
    })
    // [( , 1100), (!, 1101), (d, 000), (e, 001), (h, 0110), (l, 10), (o, 111), (r, 0111), (w, 010)]

}

fun word_count(string: String): HashMap<Char, Int> {
    val char_arr = string.toCharArray()
    val map = HashMap<Char, Int>()
    for(char in char_arr) {
        if (map.containsKey(char))
            map.replace(char, map.get(char)!! + 1)
        else
            map.put(char, 1)
    }
    return map
}

fun build_tree(map: HashMap<Char, Int>): Node<Char, Int>? {
    val list = ArrayList<Pair<Char, Int>>()
    for ((ch, i) in map) {
        list.add(Pair<Char, Int>(ch, i))
    }

    val node_list = ArrayList<Node<Char, Int>>()
    for (pair in list) {
        val (item, cnt) = pair
        node_list.add(Node(null, item, cnt))
    }

    val comparator = object: Comparator<Node<Char, Int>> {
        override fun compare(p0: Node<Char, Int>,p1: Node<Char, Int>): Int {
            val (_, _, p0_cnt) = p0
            val (_, _, p1_cnt) = p1
            return p0_cnt-p1_cnt
        }
    }
    val pq = PriorityQueue<Node<Char,Int>>(node_list.size, comparator)
    for(item in node_list) {
        pq.add(item)
    }
    while (pq.size >= 1) {
        val smallest = pq.poll()
        val second = pq.poll() ?: return smallest
        val new_node = Node<Char, Int>(Pair(smallest, second), '\\', smallest.count + second.count)
        pq.add(new_node)
    }
    return null
}

fun traversal(node: Node<Char, Int>, string: String = ""): Collection<Pair<Char, String>> {
    fun traversal_recursive(node: Node<Char, Int>, string: String, list: MutableList<Pair<Char, String>>) {
        val (sub, item, _) = node
        if (sub == null || item==null) {
            list.add(Pair<Char, String>(item!!, string))
            return
        }
        val (first, second) = sub
        traversal_recursive(first!!, string+"0", list)
        traversal_recursive(second!!, string+"1", list)
    }
    val list = ArrayList<Pair<Char, String>>()
    traversal_recursive(node, string, list)
    return list
}

data class Node<C, E>(val sub: Pair<Node<C, E>?, Node<C, E>?>?, val item: C?, val count: E)