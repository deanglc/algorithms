myTree = ['a',
          ['b',
           ['d', [], []],
           ['e', [], []]],
          ['c', ['f', [], []]]
          ]


def binary_tree(r):
    return [r, [], []]


# 新节点 插入 左子节点
# root [a , [],[]]
# root [a , ['b',['c',[],[]],[]],[]]
def insert_left(root: list, new_branch):
    t = root.pop(1)
    if len(t) > 1:  # 判断左子节点是否为空， 空
        root.insert(1, [new_branch, t, []])
    else:
        root.insert(1, [new_branch, [], []])
    return root


def insert_right(root: list, new_branch):
    t = root.pop(2)
    if len(t) > 1:  # 判断左子节点是否为空， 空
        root.insert(2, [new_branch, [], t])
    else:
        root.insert(2, [new_branch, [], []])
    return root


def get_root_val(root):
    return root[0]


def set_root_val(root, val):
    root[0] = val


def get_left_child(root):
    return root[1]


def get_right_child(root):
    return root[2]


# r = binary_tree(3)
# print(r)
# insert_left(r, 4)
# print(r)
# insert_left(r, 5)
# print(r)
# insert_right(r, 6)
# insert_right(r, 7)
# print(r)
# I = get_left_child(r)
# print(I)
# print(r)
# set_root_val(I, 9)
# print(I)
# insert_left(I, 11)
# print(I)
# print(get_right_child(get_right_child(r)))


def build_tree():
    r = binary_tree('a')
    insert_left(r, 'b')
    insert_right(r, 'c')
    print(r)
    print(get_left_child(r))
    print(get_right_child(r))
    insert_right(get_left_child(r), 'd')
    insert_right(get_right_child(r), 'f')
    insert_left(get_right_child(r), 'e')
    print(r)


if __name__ == '__main__':
    build_tree()
