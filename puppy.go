package puppy

import (
	"github.com/BumbuKhan/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBar() string {
	return dog.WhenGrownUp(Bark())
}

func BigBark() string {
	return dog.WhenGrownUp(Barks())
}