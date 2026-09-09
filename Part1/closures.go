package main
import 
	"fmt"

func intseq() func() int {
	i := 10
	return func() int {
		i *= 10
		return i
	}
}
func main() {

	nextint := intseq()
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	
	newints := intseq()
	fmt.Println(newints())
}