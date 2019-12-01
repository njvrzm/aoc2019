package aoc

func FuelForMass(mass int) int {
	return mass/3 - 2
}

func FuelForMassAndFuel(mass int) int {
	total := 0
	// 9 is the smallest mass that requires positive fuel: 9/3 - 2 == 1
	for mass >= 9 {
		fuel := FuelForMass(mass)
		total += fuel
		mass = fuel
	}
	return total
}

type FuelCalculator func(int) int

func TotalFuelRequired(masses []int, fuelc FuelCalculator) int {
	total := 0
	for _, mass := range masses {
		total += fuelc(mass)
	}
	return total
}
