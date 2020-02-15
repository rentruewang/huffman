import java.util.*;
import java.io.*;
import java.nio.file.*;
import java.nio.charset.*;
import java.util.stream.Stream;

public class huffman {
    public static void main(String[] args) {
        try (Stream<String> content = Files.lines(Path.of(args[0]), StandardCharsets.UTF_8)) {
            content.forEach(System.out::println);
        } catch (IOException ioe) {
            System.out.println(ioe.toString());
        }
    }
}
