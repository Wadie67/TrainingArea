package main
import "fmt"

func multiply(nums ...int) {
	fmt.Print(nums, "")
	product := 1

	for _, num := range nums {
		product *= num
	}
	fmt.Println(product)

}
func main() {
	multiply(5, 20)

	nums := []int{1, 2, 3, 4}
    multiply(nums...)
}