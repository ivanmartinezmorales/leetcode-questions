import 'package:dart_numerics/dart_numerics.dart';

class BST {
  int value;
  BST left;
  BST right;
  
  BST(int value) {
    this.value = value;
  }
}

bool validateBst(BST tree) {
  return _validateBst(tree, int64MinValue, int64MaxValue);
}

/// [_validateBst] is a helper method that does all of the recursion for
/// [validateBst]
bool _validateBst(BST tree, int min, int max) {
  if (tree == null)  return null;
  if (tree.value < min || tree.value >= max) return false;
  return _validateBst(tree.left, min, tree.value) && _validateBst(tree.right, tree.value, max);
}