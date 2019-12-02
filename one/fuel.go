package aoc

var FuelForMass FuelCalculator = func(mass int) int {
	return mass/3 - 2
}

var FuelForMassAndFuel FuelCalculator = func(mass int) int {
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

func (fc FuelCalculator) Sum(masses []int) (total int) {
	for i := range masses {
		total += fc(masses[i])
	}
	return
}
