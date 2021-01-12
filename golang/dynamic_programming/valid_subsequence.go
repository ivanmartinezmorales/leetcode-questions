package dynamic

func IsValidSubsequence(array []int, sequence []int) bool {
	arrayIndex := 0
	sequenceIndex := 0
	for arrayIndex < len(array) && sequenceIndex < len(sequence) {
		if array[arrayIndex] == sequence[sequenceIndex] {
			sequenceIndex += 1
		}
		arrayIndex += 1
	}
	return sequenceIndex == len(sequence)
}
