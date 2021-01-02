import java.util.*;

class DuplicateValue {

    public int firstDuplicateValue(int[] array) {
        // Given an array of integers between 1 and n, where n is the length of the array,
        // Write a function that returns the first interger that appears more than once in the array.
        //
        // Brute force:
        // Use hash table
        //
        // This uses O(n) time and O(n) space
        var hasDuplicate = new HashMap<Int, Boolean>();
        for (int n : array) {
            if (hasDuplicate.containsKey(n)) {
                return n;
            }
            hasDuplicate.put(n, false);
        }

        return -1;

    }

    public int findDuplicate(int[] array) {
        // For this one, use a hashset instead. There is no need to use a map.
        // O(n) time | O(n) space
        //
        HashSet<Integer> seen = new HashSet<Integer>();
        for (int n : array) {
            if (seen.contains(n)) {
                return n;
            }
                seen.add(n);
        }

        return -1;
    }
}
