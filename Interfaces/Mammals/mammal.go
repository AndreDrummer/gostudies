package mammals

type Mammal struct {
	Specie    string
	Quadruped bool
}

type MammalActions interface {
	Eat()
	Sleep()
}
