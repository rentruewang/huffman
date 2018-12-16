'''
The module that includes everything.
You have to install textract first before you use the module.
'''
import os
import struct
import sys

import textract

os.system('python3 -m pip install --upgrade textract')
script, src, txt = sys.argv
text = textract.process(src)
text = text.decode('utf-8')
text = text.replace('\x0c', ' ')
text = text.replace('\n', ' ')
text = text.replace('\t', '  ')
text = text.replace('\r', ' ')
text = text.replace('\\', '')
text = text.replace('*', '')
text = text.replace('  ', ' ')
with open(txt, 'w+') as nov:
    nov.write(text)


class WC:
    count = 1

    def __init__(self, word):
        self.word = word

    def add(self):
        self.count += 1

    def get_word(self):
        return self.word

    def get_count(self):
        return self.count


est = []
full_est = []
with open(txt) as file:
    for line in file:
        for ch in line:
            if ch not in est:
                est.append(ch)
                full_est.append(WC(ch))
            elif ch in est:
                full_est[est.index(ch)].add()
full_est.append(WC('EOF'))

total = 0
s_list = sorted(full_est, key=lambda name: name.count, reverse=True)
for i in s_list:
    print(i.word, i.count, end='\n')
    total += i.count
print(total)
del est


class node:
    left = None
    right = None
    element = None
    parent = None
    path = None

    def __init__(self, ele, par, left, right):
        self.left = left
        self.right = right
        self.element = ele
        self.parent = par
        if ele == None:
            self.times = left.get_times()+right.get_times()
        else:
            self.times = ele.get_count()

    def has_child(self):
        return (not self.left == None) and (not self.right == None)

    def get_element(self):
        return self.element

    def get_parent(self):
        return self.parent

    def get_left_child(self):
        return self.left

    def get_right_child(self):
        return self.right

    def set_parent(self, par):
        self.parent = par

    def set_left(self, left):
        self.left = left

    def set_right(self, right):
        self.right = right

    def get_times(self):
        return self.times

    def set_path(self, code):
        self.path = code

    def is_leaf(self):
        return not self.has_child()

    def get_path(self):
        return self.path


def merge(node1, node2):
    node_p = node(None, None, node1, node2)
    node1.set_parent(node_p)
    node2.set_parent(node_p)
    return node_p


node_est = []
for item in s_list:
    node_est.append(node(item, None, None, None))


def up_stream():
    last = len(node_est) - 1
    while(node_est[last].get_times() >= node_est[last-1].get_times()):
        if(last < 1):
            break
        node_est[last], node_est[last-1] = node_est[last-1], node_est[last]
        last -= 1


def make_tree():
    while(len(node_est) > 1):
        length = len(node_est)
        node_est[length-2] = merge(node_est[length-1], node_est[length-2])
        node_est.pop()
        up_stream()


make_tree()
root = node_est[0]
root.set_path('')
code_list = []


def get_code(node):
    left = node.get_left_child()
    left.set_path(node.get_path()+'0')
    if(left.get_left_child() == None):
        code_list.append(left)
    else:
        get_code(left)
    right = node.get_right_child()
    right.set_path(node.get_path()+'1')
    if(right.get_right_child() == None):
        code_list.append(right)
    else:
        get_code(right)


get_code(root)
print('end')
print(len(code_list))

i_list = {}
for ele in sorted(code_list, key=lambda element: element.get_path()):
    i_list[ele.get_element().get_word()] = ele.get_path()
for i in i_list:
    print(i, i_list[i])
enc = 'encoded'
counter = 0
with open(txt) as file, open('cached.txt', 'w+') as cached:
    for line in file:
        for ch in line:
            cached.write(i_list[ch])
            counter += len(i_list[ch])
            counter %= 8
    cached.write(i_list['EOF'])
    counter += len(i_list['EOF'])
    counter %= 8
    if(not counter == 0):
        cached.write((8-counter)*'0')

#
#
# reading encoded file
#
#


with open(enc, 'wb+') as encoded, open('cached.txt', 'r') as cached:
    ft = ''
    while True:
        ft = cached.read(8)
        if(len(ft) == 0):
            break
        if(ft == ''):
            break
        binary = struct.pack("B", int(ft, 2))
        encoded.write(binary)


# decoded below
#
#


decoded = 'decoded.txt'
c_node = root


def tree_traversal(input):
    global c_node
    if(input == '0'):
        left = c_node.get_left_child()
        if (left == None):
            c_node = root
        c_node = c_node.get_left_child()
    if(input == '1'):
        right = c_node.get_right_child()
        if(right == None):
            c_node = root
        c_node = c_node.get_right_child()


d_cache = 'd_cached.txt'
result = 'result.txt'
with open(enc, 'rb') as binary, open(d_cache, 'w+') as wr:
    ft = ''
    length = 1
    while True:
        ft = binary.read(1)
        if(len(ft) == 0):
            break
        ft = '{0:08b}'.format(ord(ft))
        wr.write(ft)


def decode():
    with open(d_cache, 'r') as cache, open(result, 'w+') as file:
        while not c_node.get_path() == i_list['EOF']:
            tree_traversal(cache.read(1))
            if(not c_node.get_element() == None):
                if(c_node.get_element().get_word() == 'EOF'):
                    return
                file.write(c_node.get_element().get_word())


decode()
