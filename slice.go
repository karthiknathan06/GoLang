package main
import "fmt"

func main() {
	fmt.Println("learn slices")
	myslice := []string{}
	fmt.Println(myslice)
	myslice1 := []int{1,2,3,4,5,6,6}
	fmt.Println(myslice1)
	slice2 := myslice1[3:5]
	fmt.Println(slice2)
	fmt.Println("len is ",len(slice2))
	fmt.Println("cap is ",cap(slice2))
	make_slice := make([]int,3,5)
	fmt.Println(make_slice)
}
