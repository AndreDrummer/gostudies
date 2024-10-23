package inheritance

import (
	inputs "github.com/AndreDrummer/gostudies/Inheritance/Inputs"
	packageMammals "github.com/AndreDrummer/gostudies/Inheritance/Interfaces/Mammals"
	packageVehicle "github.com/AndreDrummer/gostudies/Inheritance/Interfaces/Vehicle"
	packageCars "github.com/AndreDrummer/gostudies/Inheritance/Structs/Cars"
	packageDogs "github.com/AndreDrummer/gostudies/Inheritance/Structs/Dogs"
	packagePeople "github.com/AndreDrummer/gostudies/Inheritance/Structs/People"
)

func Inheritance() {
	var people []packagePeople.Person = make([]packagePeople.Person, 0)
	var dogs = make([]packageDogs.Dog, 0)
	cars := make([]packageCars.Car, 0)

	rich := inputs.Ashley
	poor := inputs.Kwami

	pinscher := inputs.Pinscher
	chowhcow := inputs.ChowChow

	fusion := inputs.Fusion
	celta := inputs.Celta

	luxuryVehicles := make([]packageVehicle.LuxuryVehicle, 0)
	vehicles := make([]packageVehicle.Vehicle, 0)
	richs := make([]packagePeople.RichPerson, 0)
	mammals := make([]packageMammals.Mammal, 0)

	dogs = append(dogs, pinscher, chowhcow)
	people = append(people, rich, poor)
	cars = append(cars, fusion, celta)

	luxuryVehicles = append(luxuryVehicles, fusion)
	richs = append(richs, rich)

	for _, dog := range dogs {
		mammals = append(mammals, dog)
	}

	for _, person := range people {
		mammals = append(mammals, person)
	}

	for _, car := range cars {
		vehicles = append(vehicles, car)
	}

	// All mamals should Eat and Sleep
	for _, m := range mammals {
		m.Eat()
		m.Sleep()
	}

	// All Vehicles should accelarate and Break
	for _, v := range vehicles {
		v.Accelerate(10)
		v.Break(5)
	}

	// All People should speak/think
	for _, p := range people {
		p.Speak()
		p.Think()
	}

	// All Rich should travel to Disney by Yacth
	for _, r := range richs {
		r.TravelByYacht()
		r.TravelToDisney()
	}

	// All Luxury Vehicle Should Turn the AV on and warn the benchs
	for _, lv := range luxuryVehicles {
		lv.TurnOnAC()
		lv.WarmBenches()
	}
}
