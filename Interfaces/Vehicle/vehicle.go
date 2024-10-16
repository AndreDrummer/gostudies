package vehicle

type Vehicle interface {
	Accelerate(int)
	Break(int)
}

type LuxuryVehicle interface {
	Vehicle
	WarmBenches()
	TurnOnAC()
}
