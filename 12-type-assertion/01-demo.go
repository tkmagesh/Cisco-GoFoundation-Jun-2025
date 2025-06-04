package main

import "fmt"

func main() {
	// var x interface{}
	var x any

	// x = 100
	// x = "Irure exercitation esse nostrud in fugiat enim nulla ut anim dolore."
	// x = true
	// x = 99.98
	// x = struct{}{}
	fmt.Println(x)

	x = 100
	// x = "Deserunt anim sit nisi labore qui consectetur nostrud."

	// unsafe
	// y := x.(int) * 2

	// safe (assert the type of data)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type switch
	// x = 100
	// x = "Duis occaecat magna in aliqua cupidatat consectetur ipsum."
	// x = true
	// x = 99.99
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, x * 0.976 =", val*0.976)
	default:
		fmt.Println("x is of unknown type")
	}

}
