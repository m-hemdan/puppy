package main

import (
	"fmt"

	"github.com/m-hemdan/dog"
)

func main() {
	fmt.Println(dog.WhenGrownUp(Bark()))
}
func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}
func BigBark() string {

	return dog.WhenGrownUp(Bark())
}
func BigBarks() string {

	return dog.WhenGrownUp(Barks())
}
