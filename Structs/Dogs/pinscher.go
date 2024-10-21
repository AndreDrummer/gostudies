package dogs

import "fmt"

type Pinscher struct {
	Name string
	Age  int
	Size float32
}

func (p Pinscher) Eat() {
	fmt.Printf("%v is eating!\n", p.Name)
}

func (p Pinscher) Sleep() {
	fmt.Printf("%v is sleeping!\n", p.Name)
}

func (p Pinscher) Bark() {
	fmt.Printf("%v is barking!\n", p.Name)
}

func (p Pinscher) Bite() {
	fmt.Printf("%v bite!\n", p.Name)
}
