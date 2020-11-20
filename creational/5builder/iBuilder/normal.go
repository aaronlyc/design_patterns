package iBuilder

import "fmt"

type normalBuilder struct {
	windowType string
	doorType string
	floor	int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
	fmt.Printf("Normal House Window Type: %s\n", n.windowType)
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
	fmt.Printf("Normal House Door Type: %s\n", n.doorType)
}

func (n *normalBuilder) setNumFloor() {
	n.floor = 2
	fmt.Printf("Normal House Num floor: %d\n", n.floor)
}

func (n *normalBuilder) getHouse() house {
	return house{
		doorType: n.doorType,
		windowType: n.windowType,
		floor: n.floor,
	}
}

