package main

import "fmt"

func main()  {
	var title string
	var copies int
	title = "For the love of go"
	copies = 99
	fmt.Println(title)
	fmt.Println(copies)
}

package main

import "fmt"

func printTitle(title string) {
	fmt.Println("Title: ", title)
}

func main()  {
	var name string 
	var copies int 
	copies = 99
	name = "Gorge RR Martin"
	printTitle(title)
}

package main

import "fmt"

func main()  {
	//var title string
	//var copies int 
	//var inSpecialOffer bool
	//var discountPercentage int

	title := "For the love of go"
	inSpecialOffer := true
	discountPercentage := 10
	//copies = 99

	if inSpecialOffer {
		fmt.Println("Book %s is in offer: \n", title)
		fmt.Println("With a  discount of %f % off", discountPercentage)
	}
}
