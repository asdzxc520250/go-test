package test

import "fmt"

type Test struct{}

func (t *Test) Print(s string) {
	fmt.Println(s)
}
