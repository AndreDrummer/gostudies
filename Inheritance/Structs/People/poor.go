package people

import (
	"fmt"

	cars "github.com/AndreDrummer/gostudies/Inheritance/Structs/Cars"
	dogs "github.com/AndreDrummer/gostudies/Inheritance/Structs/Dogs"
)

type Poor struct {
	Car     cars.Car
	Dog     dogs.Dog
	Name    string
	Ethinic string
	Height  float32
	Gender  string
}

func (r Poor) Think() {
	fmt.Printf("I'm %v and I'm just thinking travelling by Yacht because I'm poor!\n", r.Name)
}

func (r Poor) Eat() {
	fmt.Printf("I'm %v and I'm eating!\n", r.Name)
}

func (r Poor) Sleep() {
	fmt.Printf("I'm %v and I'm sleeping!\n", r.Name)
}

func (r Poor) Speak() {
	fmt.Printf("I'm %v and I'm speaking!\n", r.Name)
}
