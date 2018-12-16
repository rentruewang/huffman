class p_binary_tree:
    class node:
        def __init__(self, _element, _parent=None, _left_child=None, _right_child=None):
            self._element = _element
            self._parent = _parent
            self._left_child = _left_child
            self._right_child = _right_child
            if(self._left_child == None and self._right_child == None):
                self._height = 1
            else:
                self._height = max(self._left_child._height,
                                   self._right_child._height)+1
            if(_element == None):
                self.times = self._left_child.times + self._right_child.times
            else:
                self.times = self._element.count

        @property
        def element(self):
            return self._element

        @element.setter
        def element(self, _element):
            self._element = _element

        @property
        def parent(self):
            return self._parent

        @parent.setter
        def parent(self, _parent):
            self._parent = _parent

        @property
        def left_child(self):
            return self._left_child

        @left_child.setter
        def left_child(self, _left_child):
            self._left_child = _left_child
            self._height = max(self._left_child._height,
                               self._right_child._height)+1

        @property
        def right_child(self):
            return self._right_child

        @right_child.setter
        def right_child(self, _right_child):
            self._right_child = _right_child
            self._height = max(self._left_child._height,
                               self._right_child._height)+1

        @property
        def has_child(self):
            if(self._right_child == None and self._left_child == None):
                return False
            return True

        @property
        def is_leaf(self):
            return not self.has_child

        @property
        def height(self):
            if(self.is_leaf):
                self._height = 1
            return self._height

    class element:
        def __init__(self, w, c):
            self._word = w
            self._count = c

        @property
        def word(self):
            return self._word

        @property
        def count(self):
            return self._count
    arr = []

    def __init__(self, dic):
        for key, value in dic.items():
            self.arr.append(self.node(self.element(key, value)))

    def merge_nodes(self, sn_arr):
        # print('merge_nodes')
        length = len(sn_arr)
        new_node = self.node(None, None, sn_arr[length-1], sn_arr[length-2])
        sn_arr[length-1]._parent = new_node
        sn_arr[length-2]._parent = new_node
        return new_node

    def build_tree(self):
        # print('build_tree')
        sn_arr = sorted(
            self.arr, key=lambda el: el.element.count, reverse=True)
        length = len(sn_arr)
        while length > 1:
            sn_arr[length-2] = self.merge_nodes(sn_arr)
            sn_arr.pop()
            length -= 1
            self.up_stream_bubbling(sn_arr)
        return sn_arr[0]

    def up_stream_bubbling(self, sn_arr):
        # print('upstream')
        index = len(sn_arr)-1
        while sn_arr[index].times > sn_arr[index-1].times:
            if(index < 1):
                return
            sn_arr[index], sn_arr[index-1] = sn_arr[index-1], sn_arr[index]
            index -= 1
        while sn_arr[index].times == sn_arr[index-1].times:
            if(index < 1):
                return
            if(sn_arr[index]._height < sn_arr[index-1]._height):
                sn_arr[index], sn_arr[index-1] = sn_arr[index-1], sn_arr[index]
                index -= 1
            else:
                break
