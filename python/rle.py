import struct
import sys
from math import *

script, file = sys.argv
with open('encoded/binary_'+file, 'rb') as f, open('encoded/rle_'+file, 'wb+') as new:
    cur = b''
    past = b''
    count = 1
    cur = f.read(1)
    binary = []
    total = 0
    while True:
        past = cur
        cur = f.read(1)
        if past == cur:
            count += 1
        else:
            binary.extend([(count-1)*'0'+'1', '{0:08b}'.format(ord(past))])
            total += count
            count = 1
        if cur == b'':
            break
    binary.append(((8-(total % 8)) % 8)*'0')
    string = ''.join(binary)
    l = len(string)
    for i in range(0, l, 8):
        new.write(struct.pack('B', int(string[i:i+8], 2)))
with open('encoded/rle_'+file, 'rb') as f, open('decoded/rle_decoded_'+file, 'w+') as de:
    list = []
    while True:
        s = f.read(1)
        if s == b'':
            break
        list.append(str('{0:08b}'.format(ord(s))))
    string = ''.join(list)
    list = []
    state = True
    # False for ch
    # True for count
    i = 0
    l = len(string)
    while i < l:
        if state:
            j = 0
            while string[i] == '0':
                i += 1
                j += 1
                if(i==l):
                    break
            i += 1
            list.append(j+1)
        else:
            list.append(chr(int(string[i:i+8], 2)))
            i += 8
        state = not state
    state = True
    num = 0
    for i in list:
        if state:
            num = i
        else:
            de.write(num*i)
        state = not state
