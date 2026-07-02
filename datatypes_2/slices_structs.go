package main

import "fmt"

type Person struct {
	Name string
	Age int
	City string
}

func main()  {
	// Slice of structs 
	people := []Person{
		{Name: "Alice", Age: 30, City: "New York"},
		{Name: "Bob", Age: 25, City: "London"},
		{Name: "Charlie", Age: 35, City: "Tokyo"},
	}

	// Add  new Person 
	people = append(people, Person{Name: "Diana", Age: 28, City: "Paris"})

	// Access and modify 
	fmt.Println(people[0].Name)
	people[1].Age = 26 

	// Itterate with index 
	for i, person := range people {
		fmt.Printf("%d: %s is %d years old from %s\n", i, person.Name, person.Age, person.City)
	}

	// Filter wth slice 
	var adults []Person
	for _, person := range people {
		if person.Age >= 30 {
			adults = append(adults, person)
		}
	}
	fmt.Printf("Adults: %+v\n", adults)
	
}
