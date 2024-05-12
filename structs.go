package main

import "fmt"

//Declare variables in order of max to min bytes
//golang compiler will create padding and allignment spaces based on the declaration order
//like string having 4 bytes should have starting address to multiple of 4 and so on
type student struct {
	a string
	b int
	c bool
}

func main(){
	fmt.Println("Learn structs")
	a := student {
			a: "karthik",
			b: 26,
			c: true,
	}
	fmt.Println(a)			
}
