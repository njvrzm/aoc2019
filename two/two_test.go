package two

import (
	"testing"

	util "github.com/njvrzm/aoc/util"
)

func TestIntcode(t *testing.T) {
	program := Program{1, 0, 0, 0, 99}
	expected := Program{2, 0, 0, 0, 99}
	RunIntcode(program)
	if program[0] != expected[0] {
		t.Errorf("Expected program[0] to be %d, was %d", expected[0], program[0])
	}
}

func TestIntcodeForInput(t *testing.T) {
	var program Program = util.NumbersFromFile("testdata/two")
	program[1] = 12
	program[2] = 2
	expected := 4945026
	RunIntcode(program)
	if program[0] != expected {
		t.Errorf("Expected program[0] to be %d, was %d", expected, program[0])
	}
}

func TestFindIntcodeSolution(t *testing.T) {
	var program Program = util.NumbersFromFile("testdata/two")
	target := 19690720
	expectedI := 52
	expectedJ := 96
	i, j := func() (int, int) {
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				altered := make(Program, len(program))
				copy(altered, program)
				altered[1] = i
				altered[2] = j
				RunIntcode(altered)
				if altered[0] == target {
					return i, j
				}
			}
		}
		return -1, -1
	}()
	if i != expectedI || j != expectedJ {
		t.Errorf("Expected values %d, %d; got %d, %d", expectedI, expectedJ, i, j)

	}

}
