import struct
import sys
from math import *

script, file = sys.argv
with open('encoded/binary_'+file, 'rb') as f, open('encoded/runlength_'+file, 'wb+') as new:
    cur = b''
    past = b''
    count = 1
    cur = f.read(1)
    while True:
        past = cur
        cur = f.read(1)
        if past == cur:
            count += 1
        else:
            new.write(struct.pack('B', int(count)))
            new.write(struct.pack('B', ord(past)))
            count = 1
        if len(cur) == 0:
            break
