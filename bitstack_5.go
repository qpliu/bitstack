package bitstack

// Alternative encoding of a stack of stack of bits without lazy evaluation.
type bitStackStack5a struct {
	stack BitStack
}

func (stackStack bitStackStack5a) size() BitStack {
	size := BitStack1b_Empty()
	stack := stackStack.stack
	for !stack.Peek().IsNil() && stack.Peek().Bit() {
		size = size.Push(true)
		stack = stack.Pop()
	}
	return size
}

func (stackStack bitStackStack5a) data() BitStack {
	stack := stackStack.stack
	if !stack.Peek().IsNil() {
		for stack.Peek().Bit() {
			stack = stack.Pop()
		}
		stack = stack.Pop()
	}
	return stack
}

func BitStackStack5a_Empty() BitStackStack {
	return bitStackStack5a{BitStack1b_Empty()}
}

func (stackStack bitStackStack5a) dataTopBits(size, data BitStack) BitStack {
	if size.Peek().IsNil() {
		return BitStack1b_Empty()
	} else if data.Peek().IsNil() {
		for !size.Peek().IsNil() {
			data = data.Push(false)
			size = size.Pop()
		}
		return data
	} else if data.Peek().Bit() {
		return stackStack.dataTopBits(size.Pop(), data.Pop().Pop()).Push(data.Pop().Peek().Bit()).Push(true)
	} else {
		return stackStack.dataTopBits(size.Pop(), data.Pop()).Push(false)
	}
}

func (stackStack bitStackStack5a) dataRestBits(size, data BitStack) BitStack {
	if data.Peek().IsNil() {
		return data
	}
	for !size.Peek().IsNil() {
		size = size.Pop()
		if data.Peek().Bit() {
			data = data.Pop().Pop()
		} else {
			data = data.Pop()
		}
	}
	return data
}

func (stackStack bitStackStack5a) dataDone(size, data BitStack) bool {
	if data.Peek().IsNil() {
		return true
	}
	for !size.Peek().IsNil() {
		size = size.Pop()
		if data.Peek().Bit() {
			return false
		} else {
			data = data.Pop()
		}
	}
	return true
}

func (stackStack bitStackStack5a) Push(stack BitStack) BitStackStack {
	size := stackStack.size()
	var push func(BitStack, BitStack) BitStack
	push = func(stack, data BitStack) BitStack {
		topBits := stackStack.dataTopBits(size, data)
		restBits := stackStack.dataRestBits(size, data)
		if stack.Peek().IsNil() {
			topBits = topBits.Push(false)
			if stackStack.dataDone(size, data) {
				return topBits
			}
		} else {
			topBits = topBits.Push(stack.Peek().Bit()).Push(true)
			stack = stack.Pop()
		}
		return Append4b(topBits, push(stack, restBits))
	}
	stackStack.stack = Append4b(size.Push(true), push(stack, stackStack.data()).Push(false))
	return stackStack
}

func (stackStack bitStackStack5a) Pop() BitStackStack {
	size := stackStack.size()
	if size.Peek().IsNil() {
		panic("Pop empty bitStackStack5a")
	}
	var pop func(data BitStack) BitStack
	pop = func(data BitStack) BitStack {
		topBits := stackStack.dataTopBits(size, data)
		if topBits.Peek().Bit() {
			topBits = topBits.Pop().Pop()
		} else {
			topBits = topBits.Pop()
		}
		restBits := stackStack.dataRestBits(size, data)
		if stackStack.dataDone(size.Pop(), topBits) {
			return topBits
		}
		return Append4b(topBits, pop(restBits))
	}
	stackStack.stack = Append4b(size.Pop(), pop(stackStack.data()).Push(false))
	return stackStack
}

func (stackStack bitStackStack5a) Peek() BitStack {
	size := stackStack.size()
	if size.Peek().IsNil() {
		return nil
	}
	var peek func(BitStack) BitStack
	peek = func(data BitStack) BitStack {
		if !data.Peek().Bit() {
			return BitStack1b_Empty()
		}
		data = data.Pop()
		bit := data.Peek().Bit()
		data = data.Pop()
		for s := size.Pop(); !s.Peek().IsNil(); s = s.Pop() {
			if data.Peek().Bit() {
				data = data.Pop().Pop()
			} else {
				data = data.Pop()
			}
		}
		return peek(data).Push(bit)
	}
	return peek(stackStack.data())
}

func (stackStack bitStackStack5a) Bits() BitStack {
	return stackStack.stack
}

// Infinite stack
type bitStack5b bool

func BitStack5b(bit bool) BitStack {
	return bitStack5b(bit)
}

func (stack bitStack5b) Push(bit bool) BitStack {
	if bool(stack) == bit {
		return stack
	}
	return Append4b(BitStack1b_Empty().Push(bit), stack)
}

func (stack bitStack5b) Pop() BitStack {
	return stack
}

func (stack bitStack5b) Peek() MaybeBit {
	return BitToMaybeBit(bool(stack))
}

// Lazy stack
type bitStack5c struct {
	eval func() (MaybeBit, BitStack)
	bit  MaybeBit
	rest BitStack
}

func BitStack5c(eval func() (MaybeBit, BitStack)) BitStack {
	return &bitStack5c{eval, MaybeBitNil, nil}
}

func BitStack5c_Empty() BitStack {
	return &bitStack5c{nil, MaybeBitNil, nil}
}

func (stack *bitStack5c) force() {
	if stack.eval != nil {
		stack.bit, stack.rest = stack.eval()
		stack.eval = nil
	}
}

func (stack *bitStack5c) Push(bit bool) BitStack {
	return &bitStack5c{nil, BitToMaybeBit(bit), stack}
}

func (stack *bitStack5c) Pop() BitStack {
	stack.force()
	if stack.bit.IsNil() {
		panic("Pop empty bitStack5c")
	}
	return stack.rest
}

func (stack *bitStack5c) Peek() MaybeBit {
	stack.force()
	return stack.bit
}

// Somewhat lazy stack of stack of bits
type bitStackStack5d struct {
	stack BitStack
}

func (stackStack bitStackStack5d) size() BitStack {
	size := BitStack1b_Empty()
	stack := stackStack.stack
	for !stack.Peek().IsNil() && stack.Peek().Bit() {
		size = size.Push(true)
		stack = stack.Pop()
	}
	return size
}

func (stackStack bitStackStack5d) data() BitStack {
	stack := stackStack.stack
	if !stack.Peek().IsNil() {
		for stack.Peek().Bit() {
			stack = stack.Pop()
		}
		stack = stack.Pop()
	}
	return stack
}

func (stackStack bitStackStack5d) dataTopBits(size, data BitStack) BitStack {
	if size.Peek().IsNil() {
		return BitStack1b_Empty()
	} else if data.Peek().IsNil() {
		for !size.Peek().IsNil() {
			data = data.Push(false)
			size = size.Pop()
		}
		return data
	} else if data.Peek().Bit() {
		return stackStack.dataTopBits(size.Pop(), data.Pop().Pop()).Push(data.Pop().Peek().Bit()).Push(true)
	} else {
		return stackStack.dataTopBits(size.Pop(), data.Pop()).Push(false)
	}
}

func (stackStack bitStackStack5d) dataRestBits(size, data BitStack) BitStack {
	if data.Peek().IsNil() {
		return data
	}
	for !size.Peek().IsNil() {
		size = size.Pop()
		if data.Peek().Bit() {
			data = data.Pop().Pop()
		} else {
			data = data.Pop()
		}
	}
	return data
}

func BitStackStack5d_Empty() BitStackStack {
	return bitStackStack5d{BitStack1b_Empty().Push(false)}
}

func (stackStack bitStackStack5d) Push(stack BitStack) BitStackStack {
	size := stackStack.size()
	var push func(BitStack, BitStack) BitStack
	push = func(stack, data BitStack) BitStack {
		topBits := stackStack.dataTopBits(size, data)
		restBits := stackStack.dataRestBits(size, data)
		if stack.Peek().IsNil() {
			topBits = topBits.Push(false)
		} else {
			topBits = topBits.Push(stack.Peek().Bit()).Push(true)
			stack = stack.Pop()
		}
		return BitStack5c(func() (MaybeBit, BitStack) {
			return topBits.Peek(), Append4b(topBits.Pop(), push(stack, restBits))
		})
	}
	stackStack.stack = Append4b(size.Push(true), push(stack, stackStack.data()).Push(false))
	return stackStack
}

func (stackStack bitStackStack5d) Pop() BitStackStack {
	size := stackStack.size()
	if size.Peek().IsNil() {
		panic("Pop empty bitStackStack5a")
	}
	var pop func(data BitStack) BitStack
	pop = func(data BitStack) BitStack {
		topBits := stackStack.dataTopBits(size, data)
		if topBits.Peek().Bit() {
			topBits = topBits.Pop().Pop()
		} else {
			topBits = topBits.Pop()
		}
		restBits := stackStack.dataRestBits(size, data)
		return BitStack5c(func() (MaybeBit, BitStack) {
			return topBits.Peek(), Append4b(topBits.Pop(), pop(restBits))
		})
		return Append4b(topBits, pop(restBits))
	}
	if size.Pop().Peek().IsNil() {
		stackStack.stack = BitStack1b_Empty()
	} else {
		stackStack.stack = Append4b(size.Pop(), pop(stackStack.data()).Push(false))
	}
	return stackStack
}

func (stackStack bitStackStack5d) Peek() BitStack {
	size := stackStack.size()
	if size.Peek().IsNil() {
		return nil
	}
	var peek func(BitStack) BitStack
	peek = func(data BitStack) BitStack {
		if !data.Peek().Bit() {
			return BitStack5c_Empty()
		}
		data = data.Pop()
		bit := data.Peek().Bit()
		data = data.Pop()
		for s := size.Pop(); !s.Peek().IsNil(); s = s.Pop() {
			if data.Peek().Bit() {
				data = data.Pop().Pop()
			} else {
				data = data.Pop()
			}
		}
		return BitStack5c(func() (MaybeBit, BitStack) {
			return BitToMaybeBit(bit), peek(data)
		})
	}
	return peek(stackStack.data())
}

func (stackStack bitStackStack5d) Bits() BitStack {
	return stackStack.stack
}
