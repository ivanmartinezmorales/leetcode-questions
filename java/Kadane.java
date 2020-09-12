import java.util.*;

public class Kadane {
		public static int kadanesAlgorithm(int[] array) {
				int localMax = array[0];
				int globalMax = array[0];
				for (int i = 1; i < array.length; i++) {
						int val = array[i];
						localMax = Math.max(val, localMax + val);
						globalMax = Math.max(localMax, globalMax);
				}
				return globalMax;
		}
}
