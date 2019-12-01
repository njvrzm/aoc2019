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
