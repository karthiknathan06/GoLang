package main
import "fmt"
func main() {
	fmt.Println("learn arrays")
	var myarr = [3]string{"karthik"}
	for i, input := range myarr {
		fmt.Println(i, input)
	}
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println("length is ", len(arr1))
  	arr2 := [...]int{1:23,3:45}
	fmt.Println(arr2)
}
