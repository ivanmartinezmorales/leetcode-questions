class FirstDuplicate {
    func firstDuplicateValue(_ array: inout [Int]) -> Int {
        var seen = Set<Int>()
        for value in array {
            if seen.contains(value) {
                return value
            } else {
                seen.insert(value)
            }
        }

        return -1
    }
}
