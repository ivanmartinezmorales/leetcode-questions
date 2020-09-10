import java.util.*;

class InverseBST {
	// Uses O(n) space and consumes O(n) time
  public static void invertBinaryTree(BinaryTree tree) {
		ArrayDeque<BinaryTree> queue = new ArrayDeque<BinaryTree>();
		queue.addLast(tree); // adding it into the queue
		while (queue.size() > 0) {
			BinaryTree current = queue.pollFirst();
			swap(current);
			if (current.left != null) {
				queue.addLast(current.left);
			}	
			if (current.right != null) {
				queue.addLast(current.right);
			}
		}
  }
	
	// Recursive method, uses O(d) space instead of O(n) space
	public static void invertBinaryTreeRecursive(BinaryTree tree) {
		if (tree == null) return;
		swap(tree);
		invertBinaryTreeRecursive(tree.left);
		invertBinaryTreeRecursive(tree.right);
	}
	
	private static void swap(BinaryTree tree) {
		BinaryTree temp = tree.left;
		tree.left = tree.right;
		tree.right = temp;
	}

  static class BinaryTree {
    public int value;
    public BinaryTree left;
    public BinaryTree right;

    public BinaryTree(int value) {
      this.value = value;
    }
  }
}
