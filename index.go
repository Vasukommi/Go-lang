package main

import (
	"fmt"
)

func main() {
	// declare and initialize a variable
	//In summary, Go supports two ways of declaring and initializing variables:
	//var x int declares a variable x of type int
	//x := 5 declares a variable x and assigns 5 to it, Go infers the type based on the value
	var z int
	z = 5
	fmt.Println(z, " - Variable declaration with type")
	var x = 5
	y := 10 // short variable declaration

	// control structures
	if x > y {
		fmt.Println("true condition statement")
	} else {
		fmt.Println("false condition statement")
	}

	//for loop
	for i := 0; i < x; i++ {
		fmt.Println(i, " - for loop")
	}

	//switch statement
	switch x {
	case 1:
		fmt.Println("switch 5")
	case 5:
		fmt.Println("switch 1")
	default:
		fmt.Println("switch 0")
	}

	//Data Types

	/* Integers: Integers are whole numbers, and Go has several types of integers, including int, int8, int16, int32, and int64.
	Each type has a specific range of values it can represent. For example, int8 can represent values from -128 to 127,  while
	int64 can represent values from -9223372036854775808 to 9223372036854775807. */

	var x1 int = 5
	var x2 int8 = 10
	var x3 int64 = 15
	fmt.Println("Integers - data types - ", x1, x2, x3)

	/* Floating-point numbers: Floating-point numbers are numbers with a decimal point, and Go has two types of floating-point
	numbers: float32 and float64. float32 has a precision of 6 decimal places, while float64 has a precision of 15 decimal places. */

	var x4 float32 = 5.34
	var x5 float64 = 5.98765
	fmt.Println("Floating-point numbers - data types - ", x4, x5)

	/* Strings: Strings are sequences of characters, and Go has a built-in string type for representing strings.
	Go strings are immutable, which means that once a string is created, its value cannot be changed. */

	var x6 string = "vasu kommi"
	x7 := "vasu"
	fmt.Println("Strings - data types - ", x6, x7)

	/* Boolean values: Boolean values are used to represent true or false,
	and Go has a built-in bool type for representing boolean values.*/

	var x8 bool = false
	x9 := true
	fmt.Println("Boolean values - ", x8, x9)

	/* nil: nil is a special value that can be assigned to variables of any type,
	similar to null in other languages. It is used to represent the absence of a value or an uninitialized variable. */

	var e error = nil
	fmt.Println("nil - ", e)

	//Calling add function
	addNumbers := add(26, 78)
	fmt.Println(addNumbers, "Function declaration")

	//Calling divide function
	divideResult, error := divide(11, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("function - divide - ", divideResult)
	}

	//Calling sumOfNumbers function
	totalSum := sumOfNumbers(1, 2, 3, 4, 5)
	fmt.Println(totalSum, "sumOfNumbers")

	/* Go, a pointer is a variable that stores the memory address of another variable.
	A pointer holds the memory address of a value stored elsewhere in memory. */
	var examplePointer = "pointer"
	pointerAddress := &examplePointer
	pointerValue := *pointerAddress
	fmt.Println("pointerAddress", pointerAddress)
	fmt.Println("pointerValue", pointerValue)

	/* In Go, an array is a collection of elements of the same type, that are stored in contiguous memory locations.
	The size of an array is fixed and must be specified when the array is declared.*/

	var myFirstArray [5]int
	myFirstArray[0] = 1
	myFirstArray[1] = 2
	myFirstArray[2] = 3
	myFirstArray[3] = 4
	myFirstArray[4] = 5

	fmt.Println(myFirstArray, "- Go Array")

	//You can also use the short form of declaring an array and initialize it with values:

	singleLinedArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Singled Lined Array", singleLinedArray)

	/* A slice, on the other hand, is a lightweight data structure that provides a view into an underlying array.
	A slice is a dynamic data structure and its size can change after it has been created.
	A slice is defined by two indices, a start and an end, that specify the portion of the array that the slice references.*/

	simpleSlice := []int{6, 7, 8, 9, 10, 11}
	fmt.Println(simpleSlice, "simpleSlice")

	testSlice := make([]int, 5)
	fmt.Println("make", testSlice)

	//Structs and Interfaces
	myCar := Car{"Toyota", "Camry", 2019}
	myCar.Start()

	carService := Car{"Toyota", "Camry", 2019}
	cabService(carService)

} // main func end ---

/* Declaring a function: To declare a function in Go, use the func keyword, followed by the name of the function,
a list of parameters, and the return type. Here's an example: */

func add(x int, y int) int {
	return x + y
}

/* Returning multiple values: Go functions can return multiple values, which can be useful for returning multiple results or
for returning an error along with a result. Here's an example: */

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("Cannont divide with zero")
	}

	return x / y, nil
}

/* Variable number of arguments : Go also allows for functions that can take a variable number of arguments of the same
type using the ... followed by the type of the argument. It creates a slice of the type passed. Here's an example: */

func sumOfNumbers(nums ...int) int {
	totalSum := 0
	for _, eachNum := range nums {
		totalSum += eachNum
	}
	return totalSum
}

//Structs and Interfaces
/* In Go, a struct is a user-defined data type that can be used to group together related data fields.
A struct is defined using the struct keyword, followed by a set of fields enclosed in curly braces. */

type Driver interface {
	picup()
	drop()
}

type Car struct {
	Make  string
	Model string
	Year  int
}

func (c Car) Start() {
	fmt.Println("The car is starting.")
}

func (c Car) picup() {
	fmt.Println("Picked up!")
}

func (c Car) drop() {
	fmt.Println("drop!")
}

func cabService(d Driver) {
	d.picup()
	d.drop()
}
