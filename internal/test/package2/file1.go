package package2

import (
	"fmt"

	"github.com/AlekSi/gocoverutil/internal/test/package1"
)

func FizzBuzz(verbose bool) {
	for i := 1; i <= 100; i++ {
		if verbose {
			fmt.Printf("%d: ", i)
		}

		m3 := (i % 3) == 0
		m5 := (i % 5) == 0
		switch {
		case m3 && m5:
			fmt.Println(package1.FizzBuzz())
		case m3:
			fmt.Println(package1.Fizz())
		case m5:
			fmt.Println(package1.Buzz())
		default:
			fmt.Println(i)
		}
	}
}

// not tested in package2
func Foo() string {
	return "Foo"
}
