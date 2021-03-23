package controllers

type Function interface {
	Run() (interface{}, error)
}

type Controller struct{}

func (c Controller) GetFunction() (Function, error) {
	return nil, nil
}
