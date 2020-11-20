package _simple_factory

import "fmt"

type Cola interface {
	Drink() string
}

type cocaCola struct{}

func (*cocaCola) Drink() string {
	return fmt.Sprintln("Drink the cacaCola")
}

type pesiCola struct{}

func (*pesiCola) Drink() string {
	return fmt.Sprintln("Drink the pesiCola")
}

func NewCola(t int) Cola {
	switch t {
	case 1:
		fmt.Println("this is a new cocaCola")
		return &cocaCola{}
	case 2:
		fmt.Println("this is a new pesiCola")
		return &pesiCola{}
	default:
		fmt.Println("this is a error")
		return nil
	}
}
