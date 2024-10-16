package main

import (
	"fmt"

	vehicle "github.com/AndreDrummer/gostudies/Interfaces/Vehicle"
	cars "github.com/AndreDrummer/gostudies/Structs/Cars"
)

func main() {
	var Fusion cars.Fusion = cars.Fusion{
		Brand:        "Ford",
		Name:         "Fusion",
		Model:        "Ecoboost 2.0 AWD",
		CurrentSpeed: 0,
		MaxSpeed:     310,
		NWheels:      4,
		HorsePower:   234,
	}

	var Vehicle vehicle.Vehicle = &Fusion

	fmt.Printf("O %v tem um %v cavalos.\n", Fusion.Name, Fusion.HorsePower)

	Vehicle.Accelerate(10)
	Vehicle.Break(0)

}
