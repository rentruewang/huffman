import sys

script, binary = sys.argv
string = []
with open('encoded/binary_'+binary, 'rb') as bi, open('input/new_'+binary+'.txt', 'w+', encoding='utf8') as file:
    while True:
        s = bi.read(1)
        if s == b'':
            break
        if s == '':
            break
        string.append(int('{0:08b}'.format(ord(s)), 2))
    for byte in string:
        file.write(chr(byte))
