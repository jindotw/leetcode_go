package main

func robHouse(houses []int) int {
	if len(houses) == 1 {
		return houses[0]
	}

	loot0 := houses[0]
	loot1 := max(houses[0], houses[1])

	for i := 2; i < len(houses); i++ {
		loot0, loot1 = loot1, max(loot0+houses[i], loot1)
	}
	return loot1
}
