package main

import "fmt"

func main() {
	var inSpecialOffer bool
	var value float64
	var discount float64
	var title string
	var copies int
	var author string
	author = "Rick Riodarn"
	title = "For the love of go"
	copies = 39
	inSpecialOffer = true
	if inSpecialOffer {
		value = 1200.00
		discount = 10
		fmt.Println(value)
		fmt.Println("Discount %f %", discount)
	}
	fmt.Println(title)
	fmt.Println(author)
	fmt.Println(copies)
}
