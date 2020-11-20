package iBuilder

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case "normal":
		return newNormalBuilder()
	case "igloo":
		return newIglooBuilder()
	}
	return nil
}
