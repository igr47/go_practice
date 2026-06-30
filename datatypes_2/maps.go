package main

import "fmt"

func main()  {
	// Method 1: Using make 
	map1 := make(map[string]int)

	// Metod 2: Map literal 
	map2 := map[string]int{
		"apple": 5,
		"banana": 3,
		"orange": 7,
	}

	// Empty map literal 
	//map3 := map[string]string{}

	// Adding/Updating elements
	map1["one"] = 1;
	map1["two"] = 2 
	map2["grape"] = 12 // Add new key
	map2["apple"] = 10 //vmlinuz

	fmt.Println(map1)
	fmt.Println(map2)

	// Retreiving elements 
	value := map1["one"]
	fmt.Println("Value:", value)

	// Check if key exists 
	value, exists := map1["three"]
	if exists {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("key not found")
	}

	// Delete elements 
	delete(map2, "banana")
	fmt.Println(map2)

	// Length 
	fmt.Println("Length:", len(map2))
}
