package bitstack

type BitStackStack interface {
	Push(BitStack) BitStackStack
	Pop() BitStackStack
	Peek() BitStack
	Bits() BitStack
}

type bitStackStack3 struct {
	stack BitStack
}

func BitStackStack3_Empty() BitStackStack {
	return bitStackStack3{BitStack1b_Empty()}
}

func (stackStack bitStackStack3) Push(stack BitStack) BitStackStack {
	stackStack.stack = stackStack.stack.Push(false)
	for {
		if stack.Peek().IsNil() {
			return stackStack
		}
		stackStack.stack = stackStack.stack.Push(stack.Peek().Bit()).Push(true)
		stack = stack.Pop()
	}
}

func (stackStack bitStackStack3) Pop() BitStackStack {
	for {
		if stackStack.stack.Peek().Bit() {
			stackStack.stack = stackStack.stack.Pop().Pop()
		} else {
			stackStack.stack = stackStack.stack.Pop()
			return stackStack
		}
	}
}

func (stackStack bitStackStack3) Peek() BitStack {
	if stackStack.stack.Peek().IsNil() {
		return nil
	}
	stack := BitStack1b_Empty()
	for {
		if !stackStack.stack.Peek().Bit() {
			return stack
		}
		stackStack.stack = stackStack.stack.Pop()
		stack = stack.Push(stackStack.stack.Peek().Bit())
		stackStack.stack = stackStack.stack.Pop()
	}
}

func (stackStack bitStackStack3) Bits() BitStack {
	return stackStack.stack
}
