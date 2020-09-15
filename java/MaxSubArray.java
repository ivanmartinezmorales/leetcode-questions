public class MaxSubArray {

  public int maxSubArray(int[] nums) {
    // This is a good candidate for a greedy algorithmic approach
    // Basically in this approach, we're going to take an algorithm
    // that basically picks a locally optimal move at each step
    // first we need the value of the current element, and then we need the value
    // of the local maximum element, and then we need the value of the global
    // maximum that
    // we've **seen so far**.

    // So first:
    // we have two integers, one to keep track of the current sum, and one to keep
    // track of the maximum sum
    int currentSum = nums[0]; // which will be set to our first element since that first element in a possible
    // subarrray
    int maxSum = nums[0]; // is the largest sum we've seen so far. This is our global maximum
    // Now lets iterate over the entire list, omitting the first element since we've
    // already accounted for it
    for (int i = 1; i < nums.length; i++) {
      currentSum = Math.max(nums[i], currentSum + nums[i]); // finding our largest local solution so far
      maxSum = Math.max(maxSum, currentSum); // finding our global maximum solution
    }
    // when we're done iterating through the entire list, we've found our biggest
    // number in maxSum

    return maxSum;

  }
}
