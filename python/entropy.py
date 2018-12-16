import os
from math import *

from collect import *
from encoder import *


class Entropy(collector):
    def __init__(self):
        super().__init__()

    def read(self):
        ent = 0.0
        ttl = 0
        list = []
        file = '__output/'+novel+'_statistics.txt'
        with open(file, 'r', encoding='utf8') as f:
            list = f.read().split(' times.\n')
        for item in list:
            if len(item) == 0:
                continue
            tup = item.split(' occurs for ')
            ttl += int(tup[1])
        for item in list:
            if len(item) == 0:
                continue
            tup = item.split(' occurs for ')
            prob = int(tup[1])/ttl
            ent += prob*log2(1/prob)
        print('entropy is', ent)
        print('optimum compression yields', ent*ttl, 'bits')
        return ent


def entropy():
    Entropy().read()


def average():
    info = os.stat('encoded/binary_'+novel).st_size
    print('the file size is', info*8, 'bits.')


def ratio():
    info = os.stat('encoded/binary_'+novel).st_size
    compressed = info*8
    # original = os.stat('input/'+novel+'.txt').st_size*8
    original = os.stat(ipt).st_size*8
    print('the compression ratio is', original/compressed)
