package core

import "github.com/louisevanderlith/husk"

type BodyLayout = int

const (
	Sedan BodyLayout = iota
	Coupe
	Hatchback
	Van
	PickupTruck
	StationWagon
	Convertible
	SUV
	Fastback
)

type Body struct {
	Code      string
	Layout    string
	Doors     int
	StartYear int
	EndYear   int
}

func (m Body) Valid() error {
	return husk.ValidateStruct(m)
}
