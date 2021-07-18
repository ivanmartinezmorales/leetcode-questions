class TwoDimensionalArray {
    public static void main(String args[]) {
        int[][] arr = new int [10][];
        for (int i = 0; i < 10; i++) {
            arr[i] = new int[10];
            for (int j = 0; j < 10; j++) {
                arr[i][j] = i + j;
                System.out.println(arr[i][j]);
            }
        }
    }
}
