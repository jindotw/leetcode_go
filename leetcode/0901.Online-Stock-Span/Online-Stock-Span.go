package main

import "fmt"

type StockData struct {
	price int
	days  int
}

type StockSpanner struct {
	st []StockData
}

func Constructor() StockSpanner {
	return StockSpanner{
		st: make([]StockData, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	days := 1
	for len(this.st) > 0 && price >= this.st[len(this.st)-1].price {
		days += this.st[len(this.st)-1].days
		this.st = this.st[:len(this.st)-1]
	}

	data := StockData{
		price: price,
		days:  days,
	}
	this.st = append(this.st, data)

	return days
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100)) // 1
	fmt.Println(obj.Next(80))  // 1
	fmt.Println(obj.Next(60))  // 1
	fmt.Println(obj.Next(70))  // 2
	fmt.Println(obj.Next(60))  // 1
	fmt.Println(obj.Next(75))  // 1 + 1 + 2 (
	fmt.Println(obj.Next(85))
}
