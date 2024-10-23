package inputs

import (
	vehicle "github.com/AndreDrummer/gostudies/Inheritance/Interfaces/Vehicle"
	cars "github.com/AndreDrummer/gostudies/Inheritance/Structs/Cars"
)

var Fusion vehicle.LuxuryVehicle = &cars.Fusion{
	Brand:        "Ford",
	Name:         "Fusion",
	Model:        "Ecoboost 2.0 AWD",
	CurrentSpeed: 0,
	MaxSpeed:     310,
	NWheels:      4,
	HorsePower:   234,
	Luxury:       true,
}

var Celta vehicle.Vehicle = &cars.Celta{
	Brand:        "Chevrolet",
	Name:         "Celta",
	Model:        "1.0",
	CurrentSpeed: 0,
	MaxSpeed:     150,
	NWheels:      4,
	HorsePower:   118,
	Luxury:       false,
}
