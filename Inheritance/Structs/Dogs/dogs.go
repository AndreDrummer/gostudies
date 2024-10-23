package dogs

import mammals "github.com/AndreDrummer/gostudies/Inheritance/Interfaces/Mammals"

type Dog interface {
	mammals.Mammal
	Bark()
	Bite()
}
