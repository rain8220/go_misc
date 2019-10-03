package main

import (
	"fmt"
	"github.com/rain8220/stringutil"
)

func Bar() string {
	return "bar"
}

func Foo() string {
	return "foo"
}

func Qux(v string) string {
	if v == "foo" {
		return Foo()
	}

	if v == "bar" {
		return Bar()
	}

	return "INVALID"
}

func main() {
	hello := "first go really wonderful"

	fmt.Println(hello)
	fmt.Println(stringutil.Reverse(hello))
	fmt.Printf(Bar())

}
