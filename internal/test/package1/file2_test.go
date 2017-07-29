package package1_test

import (
	"testing"

	. "github.com/AlekSi/gocoverutil/internal/test/package1"
	"github.com/AlekSi/gocoverutil/internal/test/package2"
)

func TestBuzz(t *testing.T) {
	if Buzz() != "Buzz" {
		t.Error("not a Buzz")
	}
}

func TestFoo(t *testing.T) {
	if package2.Foo() != "Foo" {
		t.Error("not a Foo")
	}
}
