package modules

// Math utilities module

// Average calculates the average of a slice of integers
func Average(nums []int) float64 {
    total := 0
    for _, n := range nums {
        total += n
    }
    return float64(total) / float64(len(nums))
}

// Max returns the maximum value in a slice of integers
func Max(nums []int) int {
    max := nums[0]
    for _, n := range nums {
        if n > max {
            max = n
        }
    }
    return max
}
