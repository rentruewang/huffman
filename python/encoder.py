import struct
import sys

from code_gen import *
from collect import *
from tree_builder import *

enc = 'encoded/binary_'+novel


def read_codes():
    lst = []
    dictionary = {}
    with open('./'+output+'/'+codes, 'r', encoding='utf8') as opened:
        lst = opened.read().split(',__\n')
    for i in lst:
        if(len(i) == 0):
            continue
        tup = i.split(' has code: ')
        dictionary[tup[0]] = tup[1]
    return dictionary


def buffered_file():
    dictionary = read_codes()
    with open(ipt, 'r', encoding='utf8') as txt_file:
        counter = 0
        buffer = []
        string = ''
        for line in txt_file:
            for char in line:
                counter += len(dictionary[char])
                counter %= 8
                buffer.append(dictionary[char])
        counter += len(dictionary['EOF'])
        counter %= 8
        buffer.append(dictionary['EOF'])
        remains = (8-counter) % 8
        buffer.append(remains*'0')
        string = ''.join(buffer)
        return string


def make_binary_file():
    string = buffered_file()
    l = len(string)
    with open(enc, 'wb+') as en:
        for i in range(0, l, 8):
            binary = struct.pack('B', int(string[i:i+8], 2))
            en.write(binary)
