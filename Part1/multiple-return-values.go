package main
import 
	"fmt"

func vals() (string, string, string) {
	return "michael", "charles", "bato"

}
func main() {
	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	e, _, g := vals()
	fmt.Println(e)
	fmt.Println(g)
}