package gun

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	ak47, err := getGun("ak47")
	if !assert.NoError(t, err, "参数错误! ") {
		return
	}

	musket, err2 := getGun("musket")
	if !assert.NoError(t, err2, "参数错误! ") {
		return
	}

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()

	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}