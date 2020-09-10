import java.util.*;

public class ValidateBST {
  // We need a helper function here so overload this method
  public static boolean validateBst(BST tree) {
    return validateBst(tree, Integer.MIN_VALUE, Integer.MAX_VALUE);
  }
	
	public static boolean validateBst(BST tree, int min, int max) {
		if (tree == null) return true;
		if (tree.value < min || tree.value >= max) return false;
		
		return validateBst(tree.left, min, tree.value) && validateBst(tree.right, tree.value, max);
		
	}

  static class BST {
    public int value;
    public BST left;
    public BST right;

    public BST(int value) {
      this.value = value;
    }
  }
}
