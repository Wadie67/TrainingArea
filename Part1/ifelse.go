package main
import (
	"fmt"
	"time"
)
func main() {
	a := 60
	b := 7
	if a + b == 67 {
		fmt.Println("67!!!!!!!!!!!!!!!!!!!!!!!!")
	} else {
		fmt.Println("NO!!")
	}
	 switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
}