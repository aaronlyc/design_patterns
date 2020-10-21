package _singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	a := GetInstance()
	fmt.Printf("a point is %p\n", a)

	b := GetInstance()
	fmt.Printf("b point is %p\n", b)

	c := GetInstance()
	fmt.Printf("c point is %p\n", c)
}

//a point is 0x6dc188
//b point is 0x6dc188
//c point is 0x6dc188
