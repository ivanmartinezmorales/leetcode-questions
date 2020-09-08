public class ThirdMaximumNumber {
  public static void main(String[] args) {
    int[] nums = [3,2,1];
    int solution = thirdMax(nums);
    if (solution == 1) {
      System.out.println("That's the right number!");
    }
  }
  
  public int thirdMax(int[] nums) {
        Set<Integer> maxes = new HashSet<Integer>();
        for (int num : nums) {
            maxes.add(num);
            if (maxes.size() > 3) {
                maxes.remove(Collections.min(maxes));
            }
        }
        if (maxes.size() == 3) {
            return Collections.min(maxes);
        }
        return Collections.max(maxes);
    }
}
