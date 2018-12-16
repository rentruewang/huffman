import os

from collect import *
from tree_builder import *

script, novel = sys.argv
col = collector()
stats = col.read()
tree_root = p_binary_tree(stats).build_tree()
code_list = {}
codes = novel+'_codes.txt'


def depth_first_search(current, identity):
    if(current.is_leaf):
        code_list[current.element.word] = identity
        return
    depth_first_search(current.left_child, identity+'0')
    depth_first_search(current.right_child, identity+'1')


def get_all_codes():
    root = tree_root
    depth_first_search(root, '')
    print('getting codes...')
    print('codes are as follows.')
    for key, value in code_list.items():
        print(key, value)


def save_all_codes():
    with open('./'+output+'/'+codes, 'w+', encoding='utf8') as c_file:
        for key, value in code_list.items():
            c_file.write(key+" has code: "+value+',__\n')
