/**
Basically given an array of integers, return two numbers that equal to the target sum
If none such integers exist, then return an empty array
*/
class TwoSum {
    func twoSum(array: [int], targetSum: Int) -> [Int] {
        let array = array.sorted()                

        var left = 0
        var right = array.count - 1

        while left < right {
            let leftSide = array[left] let rightSide = array[right]
            
            let currentSum = leftSide + rightSide 

            if currentSum == targetSum {
                return [leftSide, rightSide]
            } else if currentSum < targetSum {
                left = left + 1
            } else if currentSum > targetSum {
                right = right - 1
            }
        }
        return []
    }
}