package main

import (
	"fmt"

	vehicle "github.com/AndreDrummer/gostudies/Interfaces/Vehicle"
	cars "github.com/AndreDrummer/gostudies/Structs/Cars"
)

func main() {
	var Fusion cars.Fusion = cars.Fusion{
		Car: cars.Car{
			Brand:        "Ford",
			Name:         "Fusion",
			Model:        "Ecoboost 2.0 AWD",
			CurrentSpeed: 0,
			MaxSpeed:     310,
			NWheels:      4,
			HorsePower:   234,
			Luxury:       true,
		},
	}

	var Celta cars.Celta = cars.Celta{
		Car: cars.Car{
			Brand:        "Chevrolet",
			Name:         "Celta",
			Model:        "1.0",
			CurrentSpeed: 0,
			MaxSpeed:     150,
			NWheels:      4,
			HorsePower:   118,
			Luxury:       false,
		},
	}

	var LuxuryVehicle vehicle.LuxuryVehicle = &Fusion
	var Vehicle vehicle.Vehicle = &Celta

	fmt.Printf("O %v tem um %v cavalos.\n", Fusion.Car.Name, Fusion.Car.HorsePower)
	LuxuryVehicle.Accelerate(25)
	LuxuryVehicle.WarmBenches()
	LuxuryVehicle.TurnOnAC()
	LuxuryVehicle.Break(10)

	fmt.Printf("\nO %v tem um %v cavalos.\n", Celta.Car.Name, Celta.Car.HorsePower)
	Vehicle.Accelerate(10)
	Vehicle.Break(5)

	// The following line would throw an error since Celta is not a luxury vehicle
	// Vehicle.WarmBenches()
	// Vehicle.TurnOnAC()

}
