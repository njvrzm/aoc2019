package aoc

func FuelForMass(mass int) int {
	return mass/3 - 2
}

func TotalFuelForMasses(masses []int) int {
	total := 0
	for _, mass := range masses {
		total += FuelForMass(mass)
	}
	return total
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

func TotalFuelForMassesAndFuels(masses []int) int {
	total := 0
	for _, mass := range masses {
		total += FuelForMassAndFuel(mass)
	}
	return total
}
