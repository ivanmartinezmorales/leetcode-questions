package graph

import "github.com/ivanmartinezmorales/leetcode-questions/golang/utils"

func KadanesAlgorithm(array []int) int {
	localMax := array[0]
	globalMax := array[0]
	for i := 1; i < len(array); i++ {
		val := array[i]
		localMax = utils.MaxInt(val, localMax+val)
		globalMax = utils.MaxInt(localMax, globalMax)
	}
	return globalMax
}
