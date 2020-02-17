# huffman-encoding
### An efficient encoding method.

Huffman encoding has proven, and is proved, to be a brilliant method to encode,

providing an upper bound to the compressed file, just _one character_ above **entropy**.

And an lower bound, of course, of **entropy** itself.

Very, very impressed with `golang`'s speed as it finished any of Book 1-7 before `JVM` even starts up. However, the problem isn't big enough to take advantage of `golang`'s superior parallel support.

Rust is even faster though.
