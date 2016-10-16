package bitstack

import (
	"testing"
)

func testBitStackStack(t *testing.T, label string, empty func() BitStackStack) {
	stackStack := empty()
	if stackStack.Peek() != nil {
		t.Errorf("%s: stack should be empty", label)
	}
	stack1 := randomStack(BitStack1b_Empty, 20)
	stack2 := randomStack(BitStack1b_Empty, 30)
	stack3 := randomStack(BitStack1b_Empty, 40)
	stackStack = stackStack.Push(stack1)
	testStacksEqual(t, label, stack1, stackStack.Peek())
	stackStack = stackStack.Pop()
	if stackStack.Peek() != nil {
		t.Errorf("%s: stack should be empty", label)
	}
	stackStack = stackStack.Push(stack1)
	testStacksEqual(t, label, stack1, stackStack.Peek())
	stackStack = stackStack.Push(stack2)
	testStacksEqual(t, label, stack2, stackStack.Peek())
	stackStack = stackStack.Push(stack3)
	testStacksEqual(t, label, stack3, stackStack.Peek())
	testStacksEqual(t, label, stack2, stackStack.Pop().Peek())
	testStacksEqual(t, label, stack1, stackStack.Pop().Pop().Peek())
}

func TestBitStackStack3(t *testing.T) {
	testBitStackStack(t, "bitStackStack3", BitStackStack3_Empty)
}
