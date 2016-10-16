package bitstack

import (
	"math/rand"
	"testing"
)

func testStack(t *testing.T, label string, empty func() BitStack, maxDepth int) {
	stack := empty()
	if !stack.Peek().IsNil() {
		t.Errorf("%s: peek empty should be nil", label)
	}
	panicked := false
	func() {
		defer func() {
			if recover() != nil {
				panicked = true
			}
		}()
		stack.Pop()
	}()
	if !panicked {
		t.Errorf("%s: pop empty should panic", label)
	}
	stack = testStackOperations(t, label, stack, maxDepth)
	if !stack.Peek().IsNil() {
		t.Errorf("%s: stack should be empty", label)
	}
	if maxDepth > 0 {
		stack := randomStack(empty, maxDepth-1)
		testStacksEqual(t, label, stack, stack.Push(true).Pop())
	}
}

func testStackOperations(t *testing.T, label string, stack BitStack, maxDepth int) BitStack {
	if maxDepth == 0 {
		return stack
	}
	top := stack.Peek()
	stack = stack.Push(false)
	if stack.Peek() != MaybeBit0 {
		t.Errorf("%s: peek should be 0", label)
	}
	stack = stack.Pop()
	if stack.Peek() != top {
		t.Errorf("%s: peek should be top", label)
	}
	stack = stack.Push(true)
	if stack.Peek() != MaybeBit1 {
		t.Errorf("%s: peek should be 1", label)
	}
	stack = stack.Pop()
	if stack.Peek() != top {
		t.Errorf("%s: peek should be top", label)
	}
	partialDepth := (maxDepth + 1) / 2
	for i := 0; i < partialDepth; i++ {
		stack = stack.Push(false)
		if stack.Peek() != MaybeBit0 {
			t.Errorf("%s: peek should be 0", label)
		}
	}
	stack = testStackOperations(t, label, stack, maxDepth-partialDepth)
	for i := 0; i < partialDepth; i++ {
		if stack.Peek() != MaybeBit0 {
			t.Errorf("%s: peek should be 0", label)
		}
		stack = stack.Pop()
	}
	if stack.Peek() != top {
		t.Errorf("%s: peek should be top", label)
	}
	for i := 0; i < partialDepth; i++ {
		stack = stack.Push(true)
		if stack.Peek() != MaybeBit1 {
			t.Errorf("%s: peek should be 1", label)
		}
	}
	stack = testStackOperations(t, label, stack, maxDepth-partialDepth)
	for i := 0; i < partialDepth; i++ {
		if stack.Peek() != MaybeBit1 {
			t.Errorf("%s: peek should be 1", label)
		}
		stack = stack.Pop()
	}
	if stack.Peek() != top {
		t.Errorf("%s: peek should be top", label)
	}
	return stack
}

func testBoundedStack(t *testing.T, label string, empty func() BitStack, maxDepth int) {
	testStack(t, label, empty, maxDepth)
	stack := empty()
	for i := 0; i < maxDepth; i++ {
		stack = stack.Push(true)
	}
	panicked := false
	func() {
		defer func() {
			if recover() != nil {
				panicked = true
			}
		}()
		stack.Push(true)
	}()
	if !panicked {
		t.Errorf("%s: push full stack should panic", label)
	}
}

func testStacksEqual(t *testing.T, label string, stack1, stack2 BitStack) {
	for {
		if stack1.Peek() != stack2.Peek() {
			t.Errorf("%s: stacks not equal", label)
		}
		if stack1.Peek().IsNil() {
			break
		}
		stack1 = stack1.Pop()
		stack2 = stack2.Pop()
	}
}

func randomStack(empty func() BitStack, size int) BitStack {
	stack := empty()
	for i := 0; i < size; i++ {
		stack = stack.Push(rand.Intn(2) == 1)
	}
	return stack
}

func TestBitStack1a(t *testing.T) {
	testStack(t, "BitStack1a", BitStack1a_Empty, 100)
	stack := randomStack(BitStack1a_Empty, 100)
	testStacksEqual(t, "BitStack1a", stack, stack)
}

func TestBitStack1b(t *testing.T) {
	testStack(t, "BitStack1b", BitStack1b_Empty, 100)
	stack := randomStack(BitStack1b_Empty, 100)
	testStacksEqual(t, "BitStack1b", stack, stack)
}
