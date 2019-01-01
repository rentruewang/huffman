package huffman
import java.io.File
import kotlin.math.*
import java.lang.*
import kotlin.text.*
/* The function for opening a file */
fun openFile(name:String):String{
    fun read(n:String):String = File(n).inputStream().readBytes().toString(Charsets.UTF_8)
    return read(name)
}
/* Mapping from character to count per character */
fun mapping(s:String):Map<Char,Int>{
    val map: MutableMap<Char,Int> = mutableMapOf()
    for (ch in openFile(s).toCharArray()){
        if (map.containsKey(ch)){
            map[ch] = map[ch]!!.plus(1)
        }else{
            map[ch]=1
        }
    }
    return map.toList().sortedBy {(_,v)->-v}.toMap()
}
/* The tree class itself */
class Tree<E>(file:String){
    val fileName:String = file
    /* The node class */
    class Node<E>constructor(ele: E?,par:Node<E>?,var l:Node<E>?,var r:Node<E>?){
        var isLeaf: Boolean = false
            get() = (left==null)&&(right==null)
            private set
        var isRoot: Boolean = false
            get() = parent==null
        var parent = par
        get() = field
        set(value){field = value}
        var element = ele
        var left = l
        var right = r
        var height:Int = if(this.isLeaf) 1 else max(this.left!!.height,this.right!!.height)
    }
    /* Each and every entry */
    class Entry<K,V>(k: K, v: V){
        var containsSomething: Boolean = true
            get() = value==null
        var value = v
        var key = k
    }
    /* Used to store all the entries, is sorted */
    var arr: MutableList<Node<Entry<Char,Int>>> = mutableListOf()
    private var ttl:Int = 0
    /* Finally, constructor for class Tree */
    init{
        for ((ke,va) in mapping(fileName)){
            ttl+=va
            arr.add(Node(Tree.Entry(ke,va),null,null,null))
        }
        arr.add(Node(Tree.Entry('\uFFFF',1),null,null,null))
        for (item in arr){
            println(item.element!!.key.toString()+" "+item.element!!.value.toString())
        }
        println("total "+ttl)
    }
    /* merging nodes from the back of the list */
    fun merge():Node<Entry<Char,Int>>{
        val size = arr.size
        val newNode = Node(null,null,arr[size-2],arr[size-1])
        arr[size-1].parent = newNode
        arr[size-2].parent = newNode
        return newNode
    }
    /* constructing a tree from a list */
    fun build():Node<Entry<Char,Int>>{
        var length = arr.size
        while (length > 1){
            arr[length-2] = merge()
            arr.removeAt(length-1)
            length--
            bubble()
        }
        return arr[0]
    }
    /* up-heap-bubbling (perlocate up) */
    fun bubble(){
        var index = arr.size-1
        if(index<1)
            return
        while(times(arr[index])>times(arr[index-1])){
            val temp = arr[index]
            arr[index] = arr[index-1]
            arr[index-1] = temp
            index--
            if(index<1)
                return
        }
        while(times(arr[index])==times(arr[index-1])){
            if(arr[index].height<arr[index-1].height){
                val temp = arr[index]
                arr[index] = arr[index-1]
                arr[index-1] = temp
                index--
                if (index<1)
                    return
            }else
                return
        }
    }
    /* How many times a node is accessed, return children's sum if not leaf */
    fun times(node: Node<Entry<Char,Int>>):Int{
        if(node.isLeaf)
            return node.element!!.value
        return times(node.left!!)+times(node.right!!)
    }
    val dict: MutableMap<Char,String> = mutableMapOf()
    /* depth first search */
    fun dfs(current:Node<Entry<Char,Int>>,id:String){
        if(current.isLeaf){
            dict[current.element!!.key] = id
        }else{
            dfs(current.left!!,id+"0")
            dfs(current.right!!,id+"1")
        }
    }
    /* The root of the tree */
    val root: Node<Entry<Char,Int>> = build()
    fun codes(){
        var cur = root
        dfs(cur,"")
        println("codes as follows")
        for ((k,v)in dict){
            println(k+" "+v)
        }
    }
    val binaryFile = "binary/"+file
    val stream = openFile(fileName)
    /* Write into a file */
    fun write(){
        val builder = StringBuilder()
        for(letter in stream.toCharArray()){
            builder.append(dict[letter])
        }
        builder.append(dict['\uFFFF'])
        val remainder = 8-builder.length%8
        for(i in 0 until remainder)
            builder.append("0")
        val len:Int = builder.length
        if(len%8!=0){
            throw Exception()
        }
        val barr = ByteArray(len/8)
        for(i in 0 until len step 8){
            val newb = StringBuilder()
            for(j in i until i+8){
                newb.append(builder[j])
            }
            barr[i/8] = newb.toString().toInt(2).toByte()
        }
        File(binaryFile).writeBytes(barr)
    }
    val decodedFile = "decoded/"+file+".txt"
    /* Read from encoded file */
    fun decode(){
        val barr = File(binaryFile).readBytes()
        val reservoir = StringBuilder()
        for(i in barr){
            val str = bString(unsigned(i))
            reservoir.append(str)
        }
        var current = root
        val sb = StringBuilder()
        for (item in reservoir){
            if(item=='0') {
                if (current.isLeaf) {
                    if(current.element!!.key=='\uffff')
                        break
                    sb.append(current.element!!.key)
                    current = root
                }
                current = current.left!!
            }else if (item=='1'){
                if(current.isLeaf){
                    if(current.element!!.key=='\uffff')
                        break
                    sb.append((current.element!!.key))
                    current = root
                }
                current = current.right!!
            }
        }
        File(decodedFile).writeText(sb.toString())
    }
    /* A convenience class */
    fun bString(v:Int):String{
        val s = CharArray(8)
        var j = 7
        var i = v
        while(i!=0){
            s[j] = (i%2+'0'.toInt()).toChar()
            i/=2
            j--
        }
        while(j>=0){
            s[j] = '0'
            j--
        }
        return s.joinToString(separator = "")
    }
    fun unsigned(a:Byte):Int = a.toInt() and 0xFF
    /* Calculating the entropy */
    fun entropy(){
        var ex = 0.0
        for ((_,value)in mapping(fileName)){
            val p:Double = value.toDouble()/ttl.toDouble()
            ex += p* log2(1/p)
        }
        println("the entropy is "+ex)
        println("optimum compression gives "+ex.toDouble()*ttl.toDouble()+" bits.")
        val bits = File(binaryFile).length()*8
        val original = File(fileName).length()*8
        println("the actual file size is "+bits+" bits.")
        println("the compression ratio is "+original.toDouble()/bits.toDouble())
    }
}
