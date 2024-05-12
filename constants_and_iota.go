package main

import "fmt"

func main() {
	fmt.Println("learn constants")
	
	const a = 123
	fmt.Println(a)
	
	//size would be 256 bytes by default
	const b int = 1234
	fmt.Println(b)
	
	const (
		A1= iota+1
		B1 
		C1
	)
	
	fmt.Println(A1,B1,C1)

}
