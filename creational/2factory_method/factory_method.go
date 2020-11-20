package _factory_method

import "fmt"

type factory interface {
	createCola() product
}

type product interface {
	drink() string
}

type cocaCola struct {
}

type pesiCola struct {
}

func (*cocaCola) createCola() product {
	fmt.Println("the factory is cocaCola factory")
	return &cocaCola{}
}

func (*cocaCola) drink() string {
	return "drink the cocaCola"
}

func (*pesiCola) createCola() product {
	fmt.Println("the factory is cocaCola pesiCola")
	return &pesiCola{}
}

func (*pesiCola) drink() string {
	return "drink the pesiCola"
}
