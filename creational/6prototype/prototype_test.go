package _prototype

import (
	"fmt"
	"testing"
)

var manger *PrototypeManger

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manger.Get("t1")
	t2 := t1.Clone()
	fmt.Printf("t1 point is %p, t2 point is %p\n", t1, t2)
	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManger(t *testing.T) {
	c := manger.Get("t1").Clone()

	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}
}

func init() {
	manger = NewPrototypeManger()

	t1 := &Type1{
		name: "type1",
	}
	manger.Set("t1", t1)
}
