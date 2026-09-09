package main
import "fmt"
func main() {
    for n := range 100 {
        if n%10 == 0 {
            continue
        }
        fmt.Println(n)
    }
	
	i := 2
    for i <= 10 {
        fmt.Println(i)
        i = i + 2
    }
}