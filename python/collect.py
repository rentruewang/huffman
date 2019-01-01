import os
import sys

output = '__output'
script, novel = sys.argv
# ipt = 'input/'+novel+'.txt'
ipt = novel
(*DATA_DIR, novel) = novel.split('/')
DATA_DIR = os.path.expanduser(os.path.join(*DATA_DIR))
novel = novel.replace('.txt', '')
if not os.path.exists(output):
    os.makedirs(output)
# if not os.path.exists('./input/'):
#     os.makedirs('./input/')
if not os.path.exists('./encoded/'):
    os.makedirs('./encoded/')
if not os.path.exists('./decoded/'):
    os.makedirs('./decoded/')


class collector:
    stats_name = r'./'+output+'/'+novel+'_statistics.txt'

    def __init__(self):
        dict = {'EOF': 1}
        with open(ipt, 'r', encoding='utf8') as book:
            for line in book:
                for char in line:
                    if char in dict:
                        dict[char] += 1
                    else:
                        dict[char] = 1
        self.save(
            sorted(dict.items(), key=lambda x: x[1], reverse=True), self.stats_name)

    def save(self, dict, file):
        with open(file, 'w+', encoding='utf8') as f:
            for ch in dict:
                string = ch[0]+' occurs for '+str(ch[1])+' times.\n'
                f.write(string)

    def read(self):
        return self.reader(self.stats_name)

    def reader(self, file):
        ttl = 0
        d = dict()
        list = []
        with open(file, 'r', encoding='utf8') as f:
            list = f.read().split(' times.\n')
        print('These are the times each character appears.')
        for item in list:
            if len(item) == 0:
                continue
            tup = item.split(' occurs for ')
            print(tup[0], tup[1])
            ttl += int(tup[1])
            d[tup[0]] = int(tup[1])
        print('total', ttl)
        return d
