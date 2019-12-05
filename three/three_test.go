package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestTraceWire(t *testing.T) {
	wire := WireFromDescription("U3,R3")
	trace := wire.Trace()
	expectedLength := 7
	if len(trace) != expectedLength {
		t.Errorf("Expected %d places, got %d", expectedLength, len(trace))
	}
	expectedPlaces := Trace{
		Place{0, 0},
		Place{0, -3},
		Place{3, -3},
	}
	for _, place := range expectedPlaces {
		if !trace.Contains(place) {
			t.Errorf("Trace should contain Place{%d, %d} but does not", place.x, place.y)
		}
	}

}

func TestCrossing(t *testing.T) {
	testCases := []struct {
		data     string
		distance int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, testCase := range testCases {
		lines := strings.Split(testCase.data, "\n")
		one := WireFromDescription(lines[0])
		two := WireFromDescription(lines[1])
		place := ClosestCrossing(one, two)
		distance := Origin.DistanceTo(place)
		if distance != testCase.distance {
			t.Errorf("Expected distance %d, got %d", testCase.distance, distance)
		}
	}
}

func TestPartOne(t *testing.T) {
	fh, err := os.Open("testdata/three")
	defer fh.Close()
	if err != nil {
		t.Errorf(err.Error())
	}
	data := bufio.NewScanner(fh)
	data.Scan()
	one := WireFromDescription(data.Text())
	data.Scan()
	two := WireFromDescription(data.Text())
	place := ClosestCrossing(one, two)
	distance := Origin.DistanceTo(place)
	fmt.Print(distance)

}
