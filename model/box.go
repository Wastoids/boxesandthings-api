package model

type Box struct {
	ID   string
	Name string
}

func (b Box) Equals(that Box) bool {
	return (b.ID == that.ID) && (b.Name == that.Name)
}
