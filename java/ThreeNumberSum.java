public class ThreeNumberSum {
  public static List<Integer[]> threeNumberSum(int[] array, int targetSum) {
    // Write your code here.
		// Set up our destination
		List<Integer[]> triplets = new ArrayList<Integer[]>();
		// Sort the array
		// all of these values are distinct
		Arrays.sort(array); // Sorting the array is going to take O(n log n) | O(1)
		// Now we need some pointers
		
		// we want to have at least two values left over
		for (int i = 0; i < array.length - 2; i++) {
			int left = i + 1;	// Right next to the i value
			int right = array.length - 1;
			// Now we kinda find stuff by adding it all together
			while (left < right) {
				int currentSum = array[i]	 + array[left] + array[right];
				if (currentSum == targetSum) {
					// Add this triplet to the triplet array
					triplets.add(new Integer[]{array[i], array[left], array[right]});
					left++; 
					right--;
				}	else if (currentSum < targetSum) {
						left++;	
				} else {
					right--;	
				}
			}
			
		}
		
		// Now we need to iterate over the entire array
		return triplets;	
  }
}
