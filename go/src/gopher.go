package main

import (
	"fmt" //Input output
	"errors" //Allow return errors into functions
	"math"
)

// structs are an abstraction of classes in eze's mind
type person struct {
	name string
	age int
}

func main() {

	// output basic func
	// Printf() used for format string
	fmt.Println("Hello Gopher!")

	//ints or string initialize if empty initialize with '0' or ""
	// explicit way to declare variables with data types
	var x int = 5
	var y int = 7
	var sum int = x + y

	fmt.Println(sum)

	// sugar syntax
	z := 10
	k := 10
	sugar_sum := z + k

	fmt.Println(sugar_sum)

	// conditionals
	if x < y {
		fmt.Printf("%d is smaller than %d", x, y)
		fmt.Println()
	} else if z > y{
		fmt.Printf("%d is bigger than %d", z, y)
		fmt.Println()
	} else {
		fmt.Println("Bye!")
	}

	//arrays
	//size is fix
	var a [5]int
	fmt.Println(a)

	b := [5] int {5, 4, 3, 2, 1}
	fmt.Println(b)

	//slices
	//arrays have fix size, slices not, abstraction of lists
	s := [] int {5, 4, 3, 2, 1}
	s = append(s, 0)
	fmt.Println(s)

	//maps = dictionaries
						//key	//value
	vertices := make(map[string]int)

	vertices["key_one"] = 1
	vertices["key_two"] = 2
	vertices["key_three"] = 3

	fmt.Println(vertices)

	delete(vertices, "key_two")
	fmt.Println(vertices)

	//loops
	// standard for loop
	for i := 1; i <= 5; i++ {
		fmt.Print(i)
	}

	fmt.Println()

	//also while loop
	c := 0
	for c < 5 {
		fmt.Print(c)
		c++
	}

	fmt.Println()

	// loops utilities with data structures using 'range'
	arr := [] string {"a", "b", "c", "d"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	mapp := make(map[string]int)
	mapp["key_one"] = 1
	mapp["key_two"] = 2
	mapp["key_three"] = 3

	for key, value := range mapp {
		fmt.Println("key", key, "value", value)
	}

	//calling functions
	useSqrt()
	fmt.Println()

	//using above struct
	p := person{name: "Eze", age: 19}
	fmt.Println(p)

	fmt.Println(p.name)
}

func sum (x int, y int) int {
	return x + y
}
						// this function can return two value
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Number can't be negative")
	}

	//if the number is not negativo, return the square plus nil, function need to return two values
	return math.Sqrt(x), nil
}

func useSqrt() {
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(result)
	}
}

// 'go run' runs the program
// 'go build' creates binary file, executable. Run using ./...
