package main
import 
	"fmt"

func plus(a int, b int) int{
	return a + b
}

func minus(a int, b int) int{
	return a - b
}

func divide(a int, b int) int{
	return a / b
}

func multiply(a int, b int) int{
	return a * b
}
func main() {
	ans := plus(5, 10)
	fmt.Println("5+10=", ans)

	ans = minus(5, 10)
	fmt.Println("5-10=", ans)

	ans = divide(5, 10)
	fmt.Println("5/10=", ans)

	ans = multiply(5, 10)
	fmt.Println("5*10=", ans)
}