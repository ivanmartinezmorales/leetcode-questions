class ReverseInteger {
    public static void main(String[] args) {
        int case1 = 123;
        int expected = reverse(case1);
        System.out.println(case1);
        System.out.println(expected);
    }

    public static int reverse(int x) {
        // 123
        String reversed = new StringBuilder().append(Math.abs(x)).reverse().toString(); // 321
        try {
            return (x < 0) ? Integer.parseInt(reversed) * -1 : Integer.parseInt(reversed);
        } catch (Exception e) {
            return 0;
        }
    }
}
