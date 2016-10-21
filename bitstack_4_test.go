package bitstack

import (
	"testing"
)

func testAppendStacks(t *testing.T, label string, appendStacks func(BitStack, BitStack) BitStack, stack1, stack2 BitStack) {
	stack := appendStacks(stack1, stack2)
	stack = testStackOperations(t, label, stack, 2)
	for s1 := stack1; !s1.Peek().IsNil(); {
		if stack.Peek() != s1.Peek() {
			t.Errorf("%s: peek appended stack is wrong", label)
		}
		stack = stack.Pop()
		s1 = s1.Pop()
	}
	for s2 := stack2; !s2.Peek().IsNil(); {
		if stack.Peek() != s2.Peek() {
			t.Errorf("%s: peek appended stack is wrong", label)
		}
		stack = stack.Pop()
		s2 = s2.Pop()
	}
	if !stack.Peek().IsNil() {
		t.Errorf("%s: peek appended stack should be empty", label)
	}
}

func TestAppend4a(t *testing.T) {
	stack1 := randomStack(BitStack1b_Empty, 20)
	stack2 := randomStack(BitStack1b_Empty, 40)
	testAppendStacks(t, "TestAppend4a", Append4a, stack1, stack2)
}

func TestAppend4b(t *testing.T) {
	stack1 := randomStack(BitStack1b_Empty, 20)
	stack2 := randomStack(BitStack1b_Empty, 40)
	testAppendStacks(t, "TestAppend4b", Append4b, stack1, stack2)
}

func TestAppend4c(t *testing.T) {
	stack1 := randomStack(BitStack1b_Empty, 20)
	stack2 := randomStack(BitStack1b_Empty, 40)
	testAppendStacks(t, "TestAppend4c", Append4c, stack1, stack2)
}
