package iBuilder

import "fmt"

type iglooBuilder struct {
	windowType string
	doorType string
	floor int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (i *iglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
	fmt.Printf("Igloo House Window Type: %s\n", i.windowType)
}

func (i *iglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
	fmt.Printf("Igloo House Door Type: %s\n", i.doorType)
}

func (i *iglooBuilder) setNumFloor() {
	i.floor = 1
	fmt.Printf("Igloo House Num Floor: %d\n", i.floor)
}

func (i *iglooBuilder) getHouse() house {
	return house{
		doorType: i.doorType,
		windowType: i.windowType,
		floor: i.floor,
	}
}

