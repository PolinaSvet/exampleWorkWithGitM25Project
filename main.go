package main
import (
    "fmt"
    "log"
)
func main() {
    n := ""
    fmt.Print("Enter an integer: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("You have entered a number: %s\n", n)
}
