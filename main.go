package main

import (
	"fmt"

	"github.com/AndreDrummer/gostudies/Struct/structs"
)

func main() {
	var person structs.Person = structs.Person{
		structs.Carro{
			Name: "Ford Fusion",
		},
		Name: "André",
	}

	fmt.Println("O %v tem um %v", person.Name, person.Car.Name)
}
