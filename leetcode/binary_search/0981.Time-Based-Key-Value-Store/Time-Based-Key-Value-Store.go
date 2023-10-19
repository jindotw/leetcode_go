package main

import "fmt"

type Pair struct {
	tm  int
	val string
}

type TimeMap struct {
	data map[string][]*Pair
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]*Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	pair := &Pair{tm: timestamp, val: value}
	if arr, present := this.data[key]; present {
		this.data[key] = append(arr, pair)
	} else {
		this.data[key] = []*Pair{pair}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr, present := this.data[key]
	if !present || arr[0].tm > timestamp {
		return ""
	}

	lft, rgt := 0, len(arr)-1
	for lft <= rgt {
		mid := lft + ((rgt - lft) >> 1)
		if arr[mid].tm == timestamp {
			return arr[mid].val
		} else if timestamp > arr[mid].tm {
			lft = mid + 1
		} else {
			rgt = mid - 1
		}
	}
	if arr[rgt].tm < timestamp {
		return arr[rgt].val
	}
	return arr[lft].val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	tm := Constructor()
	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)
	fmt.Println(tm.Get("love", 5))
	fmt.Println(tm.Get("love", 10))
	fmt.Println(tm.Get("love", 15))
	fmt.Println(tm.Get("love", 20))
	fmt.Println(tm.Get("love", 25))
}
