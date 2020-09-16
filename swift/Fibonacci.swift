func Fibonacci_recursive(n: Int) -> Int {
    if n == 2 {
        return 1
    } else if n == 1 {
        return 0
    } else {
        return Fibonacci_recursive(n: n - 1) +  Fibonacci_recursive(n: n - 2)
    }
}

func Fibonacci_Dynamic(n: Int) -> Int {
    var lastTwo = [0, 1]
    var counter: Int = 3
    while counter <= n {
        let new_value = lastTwo[0] + lastTwo[1]
        lastTwo[0] = lastTwo[1]
        lastTwo[1] = new_value

        counter += 1
    }
    return n > 1 ? lastTwo[1] : lastTwo[0]
}