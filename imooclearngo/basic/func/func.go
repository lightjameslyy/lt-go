package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// func div(a, b int) (int, int) {	// return values without names
func div(a, b int) (q, r int) { // return values with names
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d):\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// pass by value, doesn't work
// func swap(a, b int) {
// b, a = a, b
// }

func swap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	q, _ := div(13, 5)
	fmt.Println(q)
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
