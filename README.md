# Huffman

### An efficient encoding method.

Huffman encoding is an optimal, greedy algorithm over a random sequence. It can guarentee an upper bound on the length of the encoded output to be as much as **1 + entropy** of the original text sequence. All with a minimal runtime. It is also a lossless compression method as well. It does so by applying a tranformation of the original text sequence to make tokens appearing more often take up less memory, while allowing tokens that appear relatively sparse taking up more space.

For example, for the sequence _'Huffffman, An Efffficient Encoding Method.'_, we already know that _f_ appears way more than other charachers. Suppose that originally all characters take 4 bits to store (there are 16 kinds of characters). If the above text sequence is the only thing we care about, what we should do is to let _f_'s take less space (potentially 3 bits or even 2 bits), while allowing _h_ for example to take up more bits. Since there are 8 _f_'s and only 1 _h_, we can save 7 bits if we shorten _f_ by 1 bit but extending _h_ by 1 bit. That's basically the idea of Huffman.

<!-- TODO: Tell stories about Huffman -->
