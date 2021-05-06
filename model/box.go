package model

import "reflect"

type Box struct {
	ID     string
	Name   string
	Things []Thing
}

func (b Box) Equals(that Box) bool {
	return (b.ID == that.ID) && (b.Name == that.Name) && (reflect.DeepEqual(b.Things, that.Things))
}
