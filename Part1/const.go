package main
import 
	"fmt"

func main() {
	const s string = "constant"
	fmt.Println(s)

	const a = 27
	const b = 40
	const c = 1

	const d = a + b * c
	const mathematical = b / a

	if d == 67{
		fmt.Println("wow 67")
	}

	fmt.Println(int64(mathematical))
}