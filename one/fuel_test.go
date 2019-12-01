package aoc

import (
	"testing"

	util "github.com/njvrzm/aoc/util"
)

func TestFuelForMass(t *testing.T) {
	var massFuels = []struct {
		mass, fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, massFuel := range massFuels {
		required := FuelForMass(massFuel.mass)
		if required != massFuel.fuel {
			t.Errorf("Fuel required for %d should be %d; got %d", massFuel.mass, massFuel.fuel, required)
		}
	}
}

func TestTotalFuelForInput(t *testing.T) {
	masses := util.NumbersFromFile("testdata/input.txt")
	total := TotalFuelForMasses(masses)
	expected := 3452245
	if total != expected {
		t.Errorf("Total fuel required should be %d", expected)
	}
}
