import java.util.*;

public class SmallestDifference {
  public static int[] smallestDifference(int[] arrayOne, int[] arrayTwo) {
    // Write your code here.
		// We're going to sort the arrays in ascending order
		// Then we're going to applying a technique that takes 
		// advantage of the properties of a sorted array to solve the question
		// assume that in a sorted array
		// if they're equal to each other, then that's the smallest difference we're going to get
		// otherwise, we can pivot through the list until we find exactly what we need
		// keep track of the difference in an array
		// Any number that comes after y5, y6, and y7 are going to be greater, and
		// we don't need to compute their diffence
		// basically we want to either decrement the second array, or increment the first array to
		// come closer
		Arrays.sort(arrayOne); // is it okay to sort the arrays in place?
		Arrays.sort(arrayTwo); // is it okay to sort the arrays in place? If not, then make a copy
		
		int idxOne = 0;
		int idxTwo = 0;
		int smallest = Integer.MAX_VALUE;
		int current = Integer.MAX_VALUE;
		int[] smallestPair = new int[2];
		
		while (idxOne < arrayOne.length && idxTwo < arrayTwo.length) {
			int firstNum = arrayOne[idxOne];
			int secondNum = arrayTwo[idxTwo];
			if (firstNum < secondNum) {
				current = secondNum - firstNum;
				// increment the left pointer
				idxOne++;
			} else if (firstNum > secondNum) {
				current = firstNum - secondNum;
				idxTwo++;
			} else {
				return new int[]{firstNum, secondNum};
			}
			
			if (smallest > current) {
				smallest = current;
				smallestPair = new int[]{firstNum, secondNum};
			}
		}
		
		return smallestPair;
  }
}
