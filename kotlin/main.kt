package huffman
fun main(args: Array<String>){
    val tree = Tree<Tree.Entry<Char,Int>>(args[0])
    tree.codes()
    tree.write()
    tree.decode()
    tree.entropy()
}
