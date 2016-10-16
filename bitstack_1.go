package bitstack

type MaybeBit byte

const (
	MaybeBitNil MaybeBit = iota
	MaybeBit0   MaybeBit = iota
	MaybeBit1   MaybeBit = iota
)

func BitToMaybeBit(bit bool) MaybeBit {
	if bit {
		return MaybeBit1
	} else {
		return MaybeBit0
	}
}

func (maybeBit MaybeBit) Bit() bool {
	if maybeBit == MaybeBit0 {
		return false
	} else if maybeBit == MaybeBit1 {
		return true
	} else {
		panic("MaybeBit is nil")
	}
}

func (maybeBit MaybeBit) IsNil() bool {
	return maybeBit != MaybeBit0 && maybeBit != MaybeBit1
}

type BitStack interface {
	Push(bit bool) BitStack
	Pop() BitStack
	Peek() MaybeBit
}

// Straightforward implementation of a stack of bits using slices
type bitStack1a []bool

func BitStack1a_Empty() BitStack {
	var empty bitStack1a
	return empty
}

func (stack bitStack1a) Push(bit bool) BitStack {
	bits := make([]bool, len(stack)+1)
	copy(bits[1:], stack)
	bits[0] = bit
	return bitStack1a(bits)
}

func (stack bitStack1a) Pop() BitStack {
	if len(stack) == 0 {
		panic("Pop empty bitStack1a")
	}
	return bitStack1a(stack[1:])
}

func (stack bitStack1a) Peek() MaybeBit {
	if len(stack) == 0 {
		return MaybeBitNil
	} else {
		return BitToMaybeBit(stack[0])
	}
}

// Straightforward implementation of a stack of bits using a linked list
type bitStack1b struct {
	top  MaybeBit
	rest BitStack
}

func BitStack1b_Empty() BitStack {
	return &bitStack1b{MaybeBitNil, nil}
}

func (stack *bitStack1b) Push(bit bool) BitStack {
	return &bitStack1b{BitToMaybeBit(bit), stack}
}

func (stack *bitStack1b) Pop() BitStack {
	if stack.top == MaybeBitNil {
		panic("Pop empty bitStack1b")
	}
	return stack.rest
}

func (stack *bitStack1b) Peek() MaybeBit {
	return stack.top
}

func BitStackToString(stack BitStack) string {
	s := ""
	for ; !stack.Peek().IsNil(); stack = stack.Pop() {
		if stack.Peek().Bit() {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}
