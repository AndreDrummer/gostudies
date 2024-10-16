package cars

import (
	"fmt"
)

type Fusion struct {
	Brand        string
	Name         string
	Model        string
	CurrentSpeed int
	MaxSpeed     int
	NWheels      int
	HorsePower   int
}

func (f *Fusion) Accelerate(speed int) {
	CurrentSpeed := f.CurrentSpeed
	f.CurrentSpeed = CurrentSpeed + speed
	fmt.Printf("Ao acelerar, a velocidade do %v era %v e agora mudou para %v.\n", f.Name, CurrentSpeed, f.CurrentSpeed)
}

func (f *Fusion) Break(speed int) {
	CurrentSpeed := f.CurrentSpeed
	f.CurrentSpeed = CurrentSpeed - speed
	fmt.Printf("Ao frear, a velocidade do %v era %v e agora mudou para %v.\n", f.Name, CurrentSpeed, f.CurrentSpeed)
}
