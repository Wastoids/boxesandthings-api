package model

type Box struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (b Box) Equals(that Box) bool {
	return (b.ID == that.ID) && (b.Name == that.Name)
}
