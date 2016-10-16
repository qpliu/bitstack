package bitstack

// Straightforward implementation of a stack of bits with
// a maximum depth of two using an array and a stack pointer
type bitStack2a struct {
	bits [2]bool
	sp   uint
}

func BitStack2a_Empty() BitStack {
	return bitStack2a{}
}

func (stack bitStack2a) Push(bit bool) BitStack {
	if stack.sp >= 2 {
		panic("Push full bitStack2a")
	}
	stack.bits[stack.sp] = bit
	stack.sp++
	return stack
}

func (stack bitStack2a) Pop() BitStack {
	if stack.sp == 0 {
		panic("Pop empty bitStack2a")
	}
	stack.sp--
	return stack
}

func (stack bitStack2a) Peek() MaybeBit {
	if stack.sp == 0 {
		return MaybeBitNil
	} else {
		return BitToMaybeBit(stack.bits[stack.sp-1])
	}
}

// Implementation of a stack of bits with a maximum depth of two
// using an array of 3 bits.
type bitStack2b [3]bool

func BitStack2b_Empty() BitStack {
	return bitStack2b{}
}

func (stack bitStack2b) Push(bit bool) BitStack {
	if stack[0] {
		panic("Push full bitStack2b")
	}
	for i := 1; i < len(stack); i++ {
		if stack[i] {
			stack[i] = bit
			stack[i-1] = true
			return stack
		}
	}
	stack[len(stack)-1] = bit
	stack[len(stack)-2] = true
	return stack
}

func (stack bitStack2b) Pop() BitStack {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] {
			stack[i] = false
			stack[i+1] = true
			return stack
		}
	}
	panic("Pop empty bitStack2b")
}

func (stack bitStack2b) Peek() MaybeBit {
	for i := 1; i < len(stack); i++ {
		if stack[i-1] {
			return BitToMaybeBit(stack[i])
		}
	}
	return MaybeBitNil
}

// Implementation of a stack of bits with a maximum depth of two
// using an enumeration.
type bitStack2c byte

const (
	bs2c_   bitStack2c = iota
	bs2c_0  bitStack2c = iota
	bs2c_1  bitStack2c = iota
	bs2c_00 bitStack2c = iota
	bs2c_01 bitStack2c = iota
	bs2c_10 bitStack2c = iota
	bs2c_11 bitStack2c = iota
)

func BitStack2c_Empty() BitStack {
	return bs2c_
}

func (stack bitStack2c) Push(bit bool) BitStack {
	if bit {
		switch stack {
		case bs2c_:
			return bs2c_1
		case bs2c_0:
			return bs2c_10
		case bs2c_1:
			return bs2c_11
		case bs2c_00, bs2c_01, bs2c_10, bs2c_11:
			panic("Push full bitStack2c")
		default:
			panic("Push invalid bitStack2c")
		}
	} else {
		switch stack {
		case bs2c_:
			return bs2c_0
		case bs2c_0:
			return bs2c_00
		case bs2c_1:
			return bs2c_01
		case bs2c_00, bs2c_01, bs2c_10, bs2c_11:
			panic("Push full bitStack2c")
		default:
			panic("Push invalid bitStack2c")
		}
	}
}

func (stack bitStack2c) Pop() BitStack {
	switch stack {
	case bs2c_:
		panic("Pop empty bitStack2c")
	case bs2c_0, bs2c_1:
		return bs2c_
	case bs2c_00, bs2c_10:
		return bs2c_0
	case bs2c_01, bs2c_11:
		return bs2c_1
	default:
		panic("Pop invalid bitStack2c")
	}
}

func (stack bitStack2c) Peek() MaybeBit {
	switch stack {
	case bs2c_:
		return MaybeBitNil
	case bs2c_0, bs2c_00, bs2c_01:
		return MaybeBit0
	case bs2c_1, bs2c_10, bs2c_11:
		return MaybeBit1
	default:
		panic("Pop invalid bitStack2c")
	}
}

// Bonus implementation of a stack of bits with a maximum depth of
// 63 using a uint64.
type bitStack2d uint64

func BitStack2d_Empty() BitStack {
	return bitStack2d(0)
}

func (stack bitStack2d) Push(bit bool) BitStack {
	if stack&0x8000000000000000 != 0 {
		panic("Push full bitStack2d")
	}
	for i := bitStack2d(0x8000000000000000); i > 1; i >>= 1 {
		if stack&(i>>1) != 0 {
			stack |= i
			if !bit {
				stack &^= i >> 1
			}
			return stack
		}
	}
	if bit {
		return bitStack2d(3)
	} else {
		return bitStack2d(2)
	}
}

func (stack bitStack2d) Pop() BitStack {
	for i := bitStack2d(0x8000000000000000); i > 1; i >>= 1 {
		if stack&i != 0 {
			if i == 2 {
				return bitStack2d(0)
			}
			stack &^= i
			stack |= i >> 1
			return stack
		}
	}
	panic("Pop empty bitStack2d")
}

func (stack bitStack2d) Peek() MaybeBit {
	for i := bitStack2d(0x8000000000000000); i > 1; i >>= 1 {
		if stack&i != 0 {
			return BitToMaybeBit(stack&(i>>1) != 0)
		}
	}
	return MaybeBitNil
}
