package package1

import (
	"fmt"
)

func Buzz() string {
	return "Buzz"
}

// not tested in package1
func FizzBuzz() string {
	return fmt.Sprintf("%s%s", Fizz(), Buzz())
}
