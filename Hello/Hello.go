package main

import (
	"fmt"

	"github.com/mhaikalkal/learnGO/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Haikal")
    fmt.Println(message)

    fmt.Println(greetings.Dayum)
}