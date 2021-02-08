class BinaryTree:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


def branchSums(root):
    sums = []
    calculateBranchSum(root, 0, sums)
    return sums

def calculateBranchSums(node, runningSum, sums):
    if not Node:
        return
    newSum = runningSum + node.value
    if node.left is None and node.Right is None:
        sums.append(newSum)
    calculateBranchSum(root.left, newSum, sums)
    calculateBranchSum(root.right, newSum, sums)
