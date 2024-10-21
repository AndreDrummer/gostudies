package dogs

import "fmt"

type ChowChow struct {
	Name string
	Age  int
	Size float32
}

func (c ChowChow) Eat() {
	fmt.Printf("%v is eating!\n", c.Name)
}

func (c ChowChow) Sleep() {
	fmt.Printf("%v is sleeping!\n", c.Name)
}

func (c ChowChow) Bark() {
	fmt.Printf("%v is barking!\n", c.Name)
}

func (c ChowChow) Bite() {
	fmt.Printf("%v bite!\n", c.Name)
}
