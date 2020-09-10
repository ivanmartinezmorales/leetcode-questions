import java.util.*

class RunningSum {
    public int[] runningSum(int[] nums) {
        List<Integer> returnList = new ArrayList<Integer>();
        int runningTotal = nums[0];
        returnList.add(runningTotal);
        
        for (int i = 1; i < nums.length; i++) {
						runningTotal += nums[i];
            returnList.add(runningTotal);
        }
        return returnList.stream().mapToInt(i -> i).toArray();    
    }
}