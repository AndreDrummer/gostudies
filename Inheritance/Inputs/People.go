package inputs

import (
	people "github.com/AndreDrummer/gostudies/Inheritance/Structs/People"
)

var Kwami people.Poor = people.Poor{
	Ethinic: "Brazilian",
	Name:    "Cleiton",
	Dog:     Pinscher,
	Gender:  "Male",
	Car:     Celta,
	Height:  1.75,
}

var Ashley people.Rich = people.Rich{
	Name:    "Ashley Long",
	Ethinic: "British",
	Gender:  "Female",
	Dog:     ChowChow,
	Car:     Fusion,
	Height:  1.65,
}
