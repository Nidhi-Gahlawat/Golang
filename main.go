package main

import (
	"fmt"

	myutil "golang/myutil"
)

// Struct example
type Person struct {
	Name    string
	Age     int
	Address string
}

// function for incrementing pointers
func increment(p *int) {
	*p++
}
func main() {
	fmt.Println(myutil.Add(2, 3))
	fmt.Println(myutil.Subtract(5, 2))

	// if-else
	var num int = 10
	if num < 0 {
		fmt.Println("The number is negative")
	} else if num > 0 && num < 10 {
		fmt.Println("The number is between 1 and 9")
	} else {
		fmt.Println("The number is greater than or equal to 10")
	}
	fmt.Println("***********************")

	// Simple for loop
	fmt.Println("Counting from 0 to 9")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Another way (similar to a while loop in C++)
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	fmt.Println("***********************")

	// break and continue
	var i int = 10
	for i < 20 {
		fmt.Printf("value of i: %d\n", i)
		i++
		if i > 15 {
			break
		}
	}
	fmt.Println("After break, Odd numbers from 0 to 9: ")
	for i = 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("***********************")

	// Switch
	var x int = 2
	switch x {
	case 1:
		fmt.Println("The number is one")
	case 2:
		fmt.Println("The number is two")
	case 3:
		fmt.Println("The number is three")
	default:
		fmt.Println("The number is not one, two or three")
	}

	fmt.Println("***********************")

	// Array
	var arr [4]int // OR arr:= []int {1,2,3,4}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	for i = 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("***********************")

	// Slice
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	slice = append(slice, 5, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(slice[3])

	fmt.Println("***********************")

	// Map
	m := make(map[string]int)
	m["apple"] = 1
	m["orange"] = 2
	m["banana"] = 3
	fmt.Println(m)
	fmt.Println(m["apple"])
	fmt.Println(m["mango"])
	delete(m, "apple")
	fmt.Println(m)

	fmt.Println("***********************")

	// Struct
	var p1 Person // one way to create a new instance of Person
	p1.Name = "Akanksha"
	p1.Age = 30
	p1.Address = "12, XYZ Apartment"
	p2 := Person{Name: "Bunty", Age: 25, Address: "Block 56, ABC street"} // another way to create a new instance of Person

	fmt.Println(p1.Name) // prints "Akanksha"
	fmt.Println(p2.Age)  // prints 25

	fmt.Println("***********************")

	// Pointers
	q := 10
	p := &q
	fmt.Println(*p) // prints 10
	increment(&i)
	fmt.Println(i) // prints 11
}
