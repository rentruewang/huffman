import sys

from code_gen import *
from collect import *
from encoder import *
from tree_builder import *

enc = 'encoded/binary_'+novel
result = 'decoded/res_'+novel+'.txt'


class tree:
    def __init__(self):
        return

    class node:
        def __init__(self, ele, par, left, rite):
            self._element = ele
            self._parent = par
            self._left = left
            self._right = rite
            if(self.is_leaf):
                self._length = self.element.path
            else:
                length = len(left.element.path)
                if(self.element == None):
                    if(not length == len(self.right.element.path)):
                        print('damn', self.right.element.path,
                              self.left.element.path)
                    self._length = length-1
                else:
                    self._length = length

        @property
        def length(self):
            return self._length

        @property
        def is_leaf(self):
            return self._left == None and self._right == None

        @property
        def element(self):
            return self._element

        @property
        def parent(self):
            return self._parent

        @property
        def left(self):
            return self._left

        @property
        def right(self):
            return self._right

        @element.setter
        def element(self, ele):
            self._element = ele

        @parent.setter
        def parent(self, par):
            self._parent = par

        @left.setter
        def left(self, left):
            self._left = left

        @right.setter
        def right(self, right):
            self._right = right

    class Element:
        def __init__(self, key, path):
            self._key = key
            self._path = path

        @property
        def key(self):
            return self._key

        @key.setter
        def key(self, char):
            self._key = char

        @property
        def path(self):
            return self._path

        @path.setter
        def path(self, string):
            self._path = string

    def merge(self, n1, n2):
        # print('merge')
        l = len(n1.element.path)
        par = self.node(self.Element(
            None, n1.element.path[:l-1]), None, n1, n2)
        n1.parent = par
        n2.parent = par
        return par

    def rebuild_tree(self):
        counter = 1
        dictionary = read_codes()
        # print('dictionary:')
        # for k,v in dictionary.items():
        #     print(k,v)
        arr = []
        for k, v in dictionary.items():
            arr.append(self.node(self.Element(k, v), None, None, None))
        while len(arr) > 1:
            counter += 1
            if counter > 200:
                break
            for i in range(len(arr)-1):
                l = len(arr[i].element.path)
                if arr[i].element.path[:l-1] == arr[i+1].element.path[:l-1] and l == len(arr[i+1].element.path):
                    if(arr[i].element.path[l-1:] < arr[i+1].element.path[l-1:]):
                        arr[i] = self.merge(arr[i], arr[i+1])
                    else:
                        arr[i] = self.merge(arr[i+1], arr[i])
                    arr.pop(i+1)
                    # print('pop')
                    break
        return arr[0]


def test():
    print('\ntree\n')
    root = tree().rebuild_tree()
    through(root, '')


def through(current, s):
    if current.is_leaf:
        print(current.element.key, current.element.path, s)
        return
    through(current.left, s+'0')
    through(current.right, s+'1')


def tree_traversal():
    root = tree().rebuild_tree()
    current = root
    dictionary = read_codes()
    lst = []
    string = ''
    with open(enc, 'rb') as bi:
        while True:
            s = bi.read(1)
            if s == b'':
                break
            lst.append(str('{0:08b}'.format(ord(s))))
        string = ''.join(lst)
    with open(result, 'w+', encoding='utf8')as re:
        for ch in string:
            if ch == '':
                break
            elif ch == '0':
                if current.is_leaf:
                    e = current.element.key
                    if e == 'EOF':
                        return
                    re.write(e)
                    current = root
                current = current.left
            elif ch == '1':
                if current.is_leaf:
                    e = current.element.key
                    if e == 'EOF':
                        return
                    re.write(e)
                    current = root
                current = current.right
