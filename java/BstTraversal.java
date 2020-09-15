import java.util.*;

class BstTraversal {
  public static List<Integer> inOrderTraverse(final BST tree, final List<Integer> array) {
		if (tree != null) {
				inOrderTraverse(tree.left, array);
				array.add(tree.value);
				inOrderTraverse(tree.right, array);
		}
		return array;
  }

  public static List<Integer> preOrderTraverse(final BST tree, final List<Integer> array) {
		if (tree != null) {
				array.add(tree.value);
				preOrderTraverse(tree.left, array);
				preOrderTraverse(tree.right, array);
		}
		return array;
  }

  public static List<Integer> postOrderTraverse(final BST tree, final List<Integer> array) {
		if (tree != null) {
				postOrderTraverse(tree.left, array);
				postOrderTraverse(tree.right, array);
				array.add(tree.value);
		}
		return array;
  }

  static class BST {
    public int value;
    public BST left;
    public BST right;

    public BST(final int value) {
      this.value = value;
    }
  }
}