package huffman;

import java.util.Optional;

public final record Node(Optional<Node> left, Optional<Node> right, char token,
        int count) {
    boolean hasLeft() {
        return left.isPresent();
    }

    boolean hasRight() {
        return right.isPresent();
    }

    boolean nonLeaf() {
        return hasLeft() || hasRight();
    }

    boolean isLeaf() {
        return !nonLeaf();
    }

    boolean hasToken() {
        return token != '\0';
    }
}
