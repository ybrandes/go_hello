package main

import (
    "fmt"
    "github.com/ybrandes/go_greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}



