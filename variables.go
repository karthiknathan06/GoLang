package main

import "fmt"

func main() {
	fmt.Println("test")
	var a int
	fmt.Printf("value is [%v]",a)
	var b float64
	fmt.Printf("\nvalue is [%v]",b)
	var c bool
	fmt.Printf("\nvalue is [%v]",c)
	var d string
	fmt.Printf("\nvalue is [%v]\n",d)
	
	e := "test"
	f := 123
	g := 1.3
	fmt.Println(e,f,g)
	
	h := float64(f)
	fmt.Println(h)
}
