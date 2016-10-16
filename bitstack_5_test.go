package bitstack

import (
	"testing"
)

func TestBitStackStack5a(t *testing.T) {
	testBitStackStack(t, "bitStackStack5a", BitStackStack5a_Empty)
}

func TestBitStack5b(t *testing.T) {
	label := "bitStack5b"
	stack := BitStack5b(true)
	for i := 0; i < 1000; i++ {
		if stack.Peek() != MaybeBit1 {
			t.Errorf("%s top of stack not 1", label)
		}
		stack = stack.Pop()
	}
	stack = stack.Push(true).Push(false).Push(true)
	if !stack.Peek().Bit() {
		t.Errorf("%s top of stack not 1", label)
	}
	stack = stack.Pop()
	if stack.Peek().Bit() {
		t.Errorf("%s top of stack not 0", label)
	}
	stack = stack.Pop()
	for i := 0; i < 2000; i++ {
		if stack.Peek() != MaybeBit1 {
			t.Errorf("%s top of stack not 1", label)
		}
		stack = stack.Pop()
	}
}

func TestBitStack5c(t *testing.T) {
	testStack(t, "bitStack5c", BitStack5c_Empty, 10)
}

func TestBitStackStack5d(t *testing.T) {
	label := "bitStackStack5d"
	testBitStackStack(t, label, BitStackStack5d_Empty)
	stackStack := BitStackStack5d_Empty()
	stack1 := randomStack(BitStack1b_Empty, 10)
	stack2 := BitStack5b(true)
	stack3 := randomStack(BitStack1b_Empty, 20)
	stackStack = stackStack.Push(stack1).Push(stack2).Push(stack3)
	testStacksEqual(t, label, stack3, stackStack.Peek())
	stackStack = stackStack.Pop()
	stack4 := stackStack.Peek()
	for i := 0; i < 2000; i++ {
		if stack4.Peek() != MaybeBit1 {
			t.Errorf("%s top of stack not 1", label)
		}
		stack4 = stack4.Pop()
	}
	stackStack = stackStack.Pop()
	testStacksEqual(t, label, stack1, stackStack.Peek())
}
