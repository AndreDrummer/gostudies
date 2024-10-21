package people

import (
	mammals "github.com/AndreDrummer/gostudies/Interfaces/Mammals"
)

type Person interface {
	mammals.Mammal
	Speak()
	Think()
}

type RichPerson interface {
	TravelByYacht()
	TravelToDisney()
}
