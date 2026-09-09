package main
import 
	"fmt"

func main() {
	var num [5]int
	
	num[1] = 5
	fmt.Println("numbers:", num[1])

	var name [5] string
	name[0] = "charles"
	name[1] = "mike"
	name[2] = "remi"
	name[3] = "rice"
	name[4] = "seym"
	fmt.Println("names:", name)
}