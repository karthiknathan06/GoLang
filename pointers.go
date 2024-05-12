package main

import "fmt"

//escape analysis for analysing whether a variable will be on stack or queue
//if we use pointer by reference it will required outside func scope, hence esc analysis determines variable inside heap else stack

func main(){
	fmt.Println("learn pointers")
	a:=10
	inc1(a)
	fmt.Println("value a in main", a)
	
	inc2(&a)
	fmt.Println("value a in main", a)
}

func inc1(a int){
	a++
	fmt.Println("value a in inc1", a)
}

func inc2(a *int){
	*a++
	fmt.Println("value a in inc2", *a)
}
