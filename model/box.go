package model

import "reflect"

type Box struct {
	Name   string
	Things []Thing
}

func (b Box) Equals(that Box) bool {
	return (b.Name == that.Name) && (reflect.DeepEqual(b.Things, that.Things))
}
