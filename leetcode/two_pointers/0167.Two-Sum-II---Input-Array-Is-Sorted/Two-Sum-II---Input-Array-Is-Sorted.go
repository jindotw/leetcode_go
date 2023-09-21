package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	find := func(num int, omit int) int {
		lft, rgt := 0, len(numbers)

		for lft < rgt {
			mid := lft + ((rgt - lft) >> 1)
			if numbers[mid] == num && mid != omit {
				return mid
			} else if num > numbers[mid] {
				lft = mid + 1
			} else {
				rgt = mid
			}
		}
		return -1
	}
	for i, num := range numbers {
		diff := target - num
		if counterpart := find(diff, i); counterpart != -1 {
			if i < counterpart {
				return []int{i + 1, counterpart + 1}
			} else {
				return []int{counterpart + 1, i + 1}
			}
		}
	}

	return nil
}

func twoSum2(numbers []int, target int) []int {
	lft, rgt := 0, len(numbers)-1

	for lft < rgt {
		sum := numbers[lft] + numbers[rgt]
		if sum == target {
			return []int{lft + 1, rgt + 1}
		}
		if sum > target {
			rgt--
		} else {
			lft++
		}
	}

	return nil
}

func main() {
	nums := []int{1, 3, 4, 4}
	target := 8
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
}
