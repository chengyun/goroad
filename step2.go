package main

import "fmt"

//struct

type Book struct {
	title string
	price float32
}

func (a Book) getPriceDiff(b Book) float32 {
	return b.price - a.price
}
func (a Book) addPrice(b Book) Book { //非指针，不改变本身值
	a.price = b.price + a.price
	return a
}
func (a *Book) addPrice2(b Book) *Book { //指针，改变原值
	a.price = b.price + a.price
	return a
}

func main() {
	bookA := Book{"English", 30.50}            //按顺序
	bookB := Book{price: 25.66, title: "热爱生命"} //制定key，可以不按顺序

	fmt.Println(bookA)
	fmt.Println(bookB)

	fmt.Println(bookA.getPriceDiff(bookB))
	bookC := bookA.addPrice(bookB)
	fmt.Println(bookA, bookB, bookC)

	bookD := bookA.addPrice2(bookB)
	fmt.Println(bookA, bookB, *bookD)
}
