class BinaryTree(object):
    def __init__(self, root_obj):
        self.key = root_obj
        self.left_child = None
        self.right_child = None

    def insert_left(self, new_node):
        if self.left_child is None:
            self.left_child = BinaryTree(new_node)
        else:
            t = BinaryTree(new_node)
            t.left_child = self.left_child
            self.left_child = t

    def insert_right(self, new_node):
        if self.right_child is None:
            self.right_child = BinaryTree(new_node)
        else:
            t = BinaryTree(new_node)
            t.right_child = self.right_child
            self.right_child = t

    def get_root_val(self):
        return self.key

    def set_root_val(self, val):
        self.key = val

    def get_left_child(self):
        return self.left_child

    def get_right_child(self):
        return self.right_child

    def pre_order(self):
        if self:
            print(self.key)
        if self.left_child:
            self.left_child.pre_order()
        if self.right_child:
            self.right_child.pre_order()


if __name__ == '__main__':
    r = BinaryTree('a')
    print(r.get_root_val())
    print(r.get_left_child())
    r.insert_left('b')
    print(r.get_left_child().get_root_val())
    r.insert_right('c')
    print(r.get_right_child().get_root_val())
    r.get_right_child().set_root_val('hello')
    print(r.get_right_child().get_root_val())

    # build tree like ['a', ['b',[],[d,[],[]]], ['c', ['e',[],[]],['f',[],[]]]]
    t = BinaryTree('a')
    t.insert_left('d')
    t.insert_left('b')

    t.insert_right('c')
    t.get_right_child().insert_left('e')
    t.get_right_child().insert_right('f')
    print(t.key)
    print(t.get_right_child().key)
    print(t.get_right_child().get_left_child().key)
    print(t.get_right_child().get_right_child().key)
