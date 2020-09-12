package main

func KadanesAlgorithm(array []int) int {
		localMax := array[0]
		globalMax := array[0]
		for i := 1; i < len(array); i++ {
				val := array[i]
				localMax = max(val, localMax + val)
				globalMax = max(localMax, globalMax)
		}

		return globalMax
}

func max(a, b int) int {
		if a > b {
				return a
		}
		return b
}
