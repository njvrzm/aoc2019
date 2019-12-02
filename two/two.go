package two

type Program []int

type OpCode struct {
	name   string
	action func(int, int) int
	stop   bool
}

func (op OpCode) Apply(p Program, index int) bool {
	if op.stop {
		return false
	}
	p[p[index+3]] = op.action(p[p[index+1]], p[p[index+2]])
	return true
}

var OpCodes map[int]OpCode = map[int]OpCode{
	1: OpCode{
		name:   "Add",
		action: func(i, j int) int { return i + j },
	},
	2: OpCode{
		name:   "Mul",
		action: func(i, j int) int { return i * j },
	},
	99: OpCode{
		name:   "Stop",
		action: nil,
		stop:   true,
	},
}

func RunIntcode(p Program) {
	for i := 0; OpCodes[p[i]].Apply(p, i); i += 4 {
	}
}
