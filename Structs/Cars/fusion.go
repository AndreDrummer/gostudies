package cars

import (
	"fmt"
)

type Fusion struct{ Car Car }

func (f *Fusion) Accelerate(speed int) {
	fusion := f.Car
	CurrentSpeed := fusion.CurrentSpeed
	fusion.CurrentSpeed = CurrentSpeed + speed
	fmt.Printf("Ao acelerar, a velocidade do %v era %v e agora mudou para %v.\n", fusion.Name, CurrentSpeed, fusion.CurrentSpeed)
}

func (f *Fusion) Break(speed int) {
	fusion := f.Car
	CurrentSpeed := fusion.CurrentSpeed
	fusion.CurrentSpeed = CurrentSpeed - speed
	fmt.Printf("Ao frear, a velocidade do %v era %v e agora mudou para %v.\n", fusion.Name, CurrentSpeed, fusion.CurrentSpeed)
}

func (f *Fusion) WarmBenches() {
	fusion := f.Car
	fmt.Printf("O %v está aquecendo os bancos...\n", fusion.Name)
}

func (f *Fusion) TurnOnAC() {
	fusion := f.Car
	fmt.Printf("O %v está com os Ar Condicionado Ligado...\n", fusion.Name)
}
