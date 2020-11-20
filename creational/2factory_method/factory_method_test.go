package _factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	//cocaCola
	cocaCola := new(cocaCola)
	cocaColaProduct := cocaCola.createCola()
	t.Log(cocaColaProduct.drink())

	//pesiCola
	pesiCola := new(pesiCola)
	pesiColaProduct := pesiCola.createCola()
	t.Log(pesiColaProduct.drink())
}
