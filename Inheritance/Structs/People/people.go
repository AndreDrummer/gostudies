package people

import (
	mammals "github.com/AndreDrummer/gostudies/Inheritance/Interfaces/Mammals"
)

type Person interface {
	mammals.Mammal
	Speak()
	Think()
}

type RichPerson interface {
	Person
	TravelByYacht()
	TravelToDisney()
}
