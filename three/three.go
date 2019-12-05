package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type Segment struct {
	direction Direction
	length    int
}

type Direction struct {
	dx int
	dy int
}

type Place struct {
	x int
	y int
}

var Origin Place = Place{0, 0}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (this Place) DistanceTo(that Place) int {
	return abs(this.x-that.x) + abs(this.y-that.y)
}

func (place Place) Towards(direction Direction) Place {
	return Place{place.x + direction.dx, place.y + direction.dy}
}

type Trace []Place

var Directions = map[string]Direction{
	"R": Direction{1, 0},
	"L": Direction{-1, 0},
	"U": Direction{0, -1},
	"D": Direction{0, 1},
}

// SegmentFromDescription reads a string of the form DNNN, where D is one
// of "R", "L", "U", or "D" (signifying right, left, up or down) and NNN is an
// integer representing a length. It returns a Segment struct matching the
// description.
func SegmentFromDescription(description string) Segment {
	direction := Directions[string(description[0])]
	length, _ := strconv.Atoi(description[1:])
	return Segment{direction, length}
}

type Wire []Segment

// WireFromDescription reads a string containing a comma-separated list of
// segment descriptions and returns a Wire (slice of segments) matching them.
func WireFromDescription(description string) Wire {
	segs := strings.Split(description, ",")
	var wire = make(Wire, len(segs))
	for i := range segs {
		wire[i] = SegmentFromDescription(segs[i])
	}
	return wire
}

func (trace Trace) Contains(target Place) bool {
	for _, place := range trace {
		if place == target {
			return true
		}
	}
	return false
}

func (wire Wire) Trace() Trace {
	where := Place{}
	trace := Trace{where}
	for _, segment := range wire {
		for i := 0; i < segment.length; i++ {
			where = where.Towards(segment.direction)
			trace = append(trace, where)
		}
	}
	return trace
}

func Crossings(w, z Wire) Trace {
	crossings := make(Trace, 0)
	wt := w.Trace()
	zt := z.Trace()
	fmt.Println(len(wt), len(zt))
	for i := range wt {
		if zt.Contains(wt[i]) {
			crossings = append(crossings, wt[i])
			fmt.Println(wt[i])
		}
	}
	return crossings
}

func ClosestCrossing(w, z Wire) Place {
	distance := 2 << 31
	closest := Place{}
	for _, crossing := range Crossings(w, z) {
		if crossing == Origin {
			continue
		}
		d := Origin.DistanceTo(crossing)
		if d < distance {
			closest = crossing
			distance = d
		}
	}
	return closest
}
