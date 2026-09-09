package main
import (
    "fmt"
    "time"
)
func main() {
	i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("Uno")
    case 2:
        fmt.Println("Dos")
    case 3:
        fmt.Println("Tres")
    }

	fmt.Println("Today is", time.Now().Weekday())
}