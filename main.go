package main
import (
    "fmt"
    "log"
)
func main() {
    n := 0
    fmt.Print("Enter an integer: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("You have entered a number: %d\n", n)
}
