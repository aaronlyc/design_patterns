package iBuilder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	fmt.Printf("window type is: %s, door type is: %s, floor nums is: %d\n", normalHouse.windowType, normalHouse.doorType, normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Printf("window type is: %s, door type is: %s, floor nums is: %d\n", iglooHouse.windowType, iglooHouse.doorType, iglooHouse.floor)
}