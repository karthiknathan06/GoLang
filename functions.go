package main

import "fmt"
import "errors"

func main() {
	fmt.Println("learn functions")
	fmt.Println("normal function ",add(2,3))
	out, err := divide(2,0)
	fmt.Println("normal function with return types ",out, err)
	addfunc := func(x float64,y float64) float64 {
		return x+y
	}
	out1 := addfunc(2.3,3.4)
	fmt.Println("anonymous function with return types ",out1)
	sum := addnumbers(1,2,3,4,5)
	fmt.Println("variable arg function with return types ",sum)
}

func add(x int,y int) int {
	return x+y
}

func divide(x float64,y float64) (float64, error) {
	if(y==0) {
		return 0, errors.New("second no is zero")
		}
	return x / y, nil
}

func addnumbers(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum =sum+num
	}
	return sum
}
