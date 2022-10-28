package huffman;

import java.util.Comparator;
import java.util.HashMap;
import java.util.Optional;
import java.util.PriorityQueue;

public class Tree {
    private Node root;

    public Tree(Node root) {
        this.root = root;
    }

    public Tree(String content) {
        this(makeRoot(content));
    }

    public HashMap<Character, String> asPath() {
        HashMap<Character, String> result = new HashMap<>();
        asPathRecursive(root, "", result);
        return result;
    }

    private static void asPathRecursive(Node node, String path, HashMap<Character, String> collection) {
        if (node.isLeaf()) {
            assert node.token() != '\0';
            collection.put(node.token(), path);
        }

        if (node.hasLeft()) {
            asPathRecursive(node.left().get(), path + "0", collection);
        }

        if (node.hasRight()) {
            asPathRecursive(node.right().get(), path + "1", collection);
        }
    }

    private static Node makeRoot(String content) {
        var array = content.toCharArray();

        HashMap<Character, Integer> wordCount = new HashMap<>();

        for (char ch : array) {
            int count = wordCount.getOrDefault(ch, 0);
            wordCount.put(ch, count + 1);
        }

        var smaller = new Comparator<Node>() {
            @Override
            public int compare(Node arg0, Node arg1) {
                return arg0.count() - arg1.count();
            }
        };

        PriorityQueue<Node> huffmanList = new PriorityQueue<>(wordCount.size(), smaller);

        for (var kv : wordCount.entrySet()) {
            char key = kv.getKey();
            int count = kv.getValue();

            var node = new Node(Optional.empty(), Optional.empty(), key, count);
            huffmanList.add(node);
        }

        for (int i = 0; i < wordCount.size() - 1; ++i) {
            Node first = huffmanList.poll();
            Node second = huffmanList.poll();

            Node merged = new Node(Optional.of(first), Optional.of(second), '\0',
                    first.count() + second.count());

            huffmanList.add(merged);
        }

        assert huffmanList.size() == 1;

        return huffmanList.poll();
    }
}
