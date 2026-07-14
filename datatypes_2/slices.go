package main

import "fmt"
// Demostrating working with slices. 

func demonstrateCapacity() {
	s := make([]int, 0, 3)
	fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)

	for i := 1; i <= 5; i++ {
		s = append(s, i)
		fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
}
func main()  {
	// Method 1: Using make 
	slice1 := make([]int,5) // length=5, capacity=5
	slice2 := make([]int, 3, 5) // length=3 capacity=5

	// Method 2: Slice literal
	slice3 := []int{1,2,3,4,5}

	// Method 3: From arra
	arr := [5]int{10,20,30,40,50}
	slice4 := arr[1:4]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// Appending elements 
	slice3 = append(slice3, 6, 7, 8)
	fmt.Println(slice3)

	// Copying slices 
	dest := make([]int, len(slice3))
	copied := copy(dest, slice3)
	fmt.Println("Copied %d elements: %v\n", copied, dest)

	// Slicing operations 
	//fmt.Println(slice3[2:5])
	//fmt.Println(slice3[:3])
	//fmt.Println(slice3[5:])

	demonstrateCapacity()
}
