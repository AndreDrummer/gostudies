package main

import (
	"fmt"

	inputs "github.com/AndreDrummer/gostudies/Inputs"
	vehicle "github.com/AndreDrummer/gostudies/Interfaces/Vehicle"
)

func main() {

	Fusion := inputs.Fusion
	var LuxuryVehicle vehicle.LuxuryVehicle = Fusion
	var Vehicle vehicle.Vehicle = LuxuryVehicle

	fmt.Println(Vehicle)
	fmt.Println(LuxuryVehicle)
	fmt.Println(Fusion)
}
