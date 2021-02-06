class Node:
    def __init__(self, name):
        self.children = []
        self.name = name

    def addChild(self, name):
        self.children.append(Node(name))
        return self

    def travereArrayDFS(self, array):
        array.append(self.name)
        for child in self.children:
            child.depthFirstSearch(array)
        return array

    def depthFirstSearch(self, val):
        if self.name == val:
            return True
        for child in self.children:
            child.depthFirstSearch(val)
        return False
