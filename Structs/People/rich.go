package people

import (
	"fmt"

	cars "github.com/AndreDrummer/gostudies/Structs/Cars"
	dogs "github.com/AndreDrummer/gostudies/Structs/Dogs"
)

type Rich struct {
	Car     cars.Car
	Dog     dogs.Dog
	Name    string
	Ethinic string
	Height  float32
	Gender  string
}

func (r Rich) TravelByYacht() {
	fmt.Printf("I'm %v and I'm travelling by Yacht because I'm rich!\n", r.Name)
}

func (r Rich) TravelToDisney() {
	fmt.Printf("I'm %v and I'm travelling to Disney because I'm rich!\n", r.Name)
}

func (r Rich) Think() {
	fmt.Printf("I'm %v and I'm thinking buying another Yacht because I'm rich!\n", r.Name)
}

func (r Rich) Eat() {
	fmt.Printf("I'm %v and I'm eating!\n", r.Name)
}

func (r Rich) Sleep() {
	fmt.Printf("I'm %v and I'm sleeping!\n", r.Name)
}

func (r Rich) Speak() {
	fmt.Printf("I'm %v and I'm speaking!\n", r.Name)
}
