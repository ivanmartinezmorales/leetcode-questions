import java.util.*;

class CountElements {
    public int countElements(int[] arr) {
        Set<Integer> set = new HashSet<Integer>();
        for (int x : arr) set.add(x);
        
        int count = 0;
        for (int x : arr) {
            if (set.contains(x + 1)) {
                count++;
            }
        }
        return count;
    }
}