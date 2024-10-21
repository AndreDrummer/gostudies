package inputs

import (
	dogs "github.com/AndreDrummer/gostudies/Structs/Dogs"
)

var Pinscher dogs.Dog = &dogs.Pinscher{
	Name: "Liu Kang",
	Age:  3,
	Size: 2,
}

var ChowChow dogs.Dog = &dogs.ChowChow{
	Name: "Chad",
	Age:  1,
	Size: 0,
}
