package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	avg := sum / 2
	memo := make([]int, avg+1)

	for _, num := range nums {
		for j := avg; j >= num; j-- {
			memo[j] = max(memo[j], memo[j-num]+num)
		}
	}

	return memo[avg] == avg
}
