import java.util.*;

class FizzBuzz {
    public List<String> fizzBuzz(int n) {
        List<String> returnList = new ArrayList<String>();
        for (int num = 1; num <= n; num++) {
            boolean div3 = (num % 3 == 0);
            boolean div5 = (num % 5 == 0);
            
            String ans = "";
            if (div3) {
                ans += "Fizz";
            }
            if (div5) {
                ans += "Buzz";
            }
            if (ans.equals("")) {
                ans += Integer.toString(num);
            }
            
            returnList.add(ans);
        }
        return returnList;
    }
}
