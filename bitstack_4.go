package bitstack

// Straightforward stack append
func Append4a(stack1, stack2 BitStack) BitStack {
	s := BitStack1b_Empty()
	for ; !stack1.Peek().IsNil(); stack1 = stack1.Pop() {
		s = s.Push(stack1.Peek().Bit())
	}
	for ; !s.Peek().IsNil(); s = s.Pop() {
		stack2 = stack2.Push(s.Peek().Bit())
	}
	return stack2
}

// Lazy stack append
type append4b [2]BitStack

func Append4b(stack1, stack2 BitStack) BitStack {
	return append4b{stack1, stack2}
}

func (stack append4b) Push(bit bool) BitStack {
	stack[0] = stack[0].Push(bit)
	return stack
}

func (stack append4b) Pop() BitStack {
	if stack[0].Peek().IsNil() {
		return stack[1].Pop()
	}
	stack[0] = stack[0].Pop()
	return stack
}

func (stack append4b) Peek() MaybeBit {
	if stack[0].Peek().IsNil() {
		return stack[1].Peek()
	}
	return stack[0].Peek()
}
