// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func Greet(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty Name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func main() {
	fmt.Println(Greet("Amol.."))
}
