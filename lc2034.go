package main

// type StockPrice struct {
// 	maxIdx, minIdx, newIdx int
// 	idx                    map[int]int
// 	price                  []int
// }

// func Constructor() StockPrice {
// 	maxIdx, minIdx, newIdx := -1, -1, -1
// 	idx := make(map[int]int, 1e5+5)
// 	price := make([]int, 0, 1e5+5)

// 	return StockPrice{maxIdx, minIdx, newIdx, idx, price}
// }

// func (this *StockPrice) Update(timestamp int, price int) {
// 	if this.idx[timestamp] != 0 {
// 		id := this.idx[timestamp]
// 		this.price[id] = price
// 	} else {
// 		this.price = append(this.price, price)
// 		this.idx[timestamp] = len(this.price) - 1
// 	}
// 	id := this.idx[timestamp]
// 	if price > this.price[this.maxIdx] {
// 		this.maxIdx = this.idx[id]
// 	}
// 	if price < this.price[this.minIdx] {
// 		this.minIdx = this.idx[id]
// 	}
// 	if timestamp >map
// }

// func (this *StockPrice) Current() int {

// }

// func (this *StockPrice) Maximum() int {

// }

// func (this *StockPrice) Minimum() int {

// }

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
