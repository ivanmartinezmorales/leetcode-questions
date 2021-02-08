class BinaryTree:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

def invert_binary_tree(root):
    queue = [root]
    while queue:
        current_node = queue.pop(0)
        if current_node:
            current_node.left, current_node.right = current_node.right, current_node.left
            queue.append(current_node.left)
            queue.append(current_node.right)

