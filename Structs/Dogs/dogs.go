package dogs

import mammals "github.com/AndreDrummer/gostudies/Interfaces/Mammals"

type Dog interface {
	mammals.Mammal
	Bark()
	Bite()
}
