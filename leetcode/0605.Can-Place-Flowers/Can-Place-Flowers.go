package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(flowerbed, 0)
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		n--
	}

	for idx := 1; n > 0 && idx < len(flowerbed)-1; idx++ {
		if flowerbed[idx] == 0 && flowerbed[idx-1] == 0 && flowerbed[idx+1] == 0 {
			flowerbed[idx] = 1
			n--
		}
	}
	return n <= 0
}

func main() {
	nums := []int{0, 0, 1, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(nums, n))
}
