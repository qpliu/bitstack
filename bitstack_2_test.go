package bitstack

import (
	"testing"
)

func TestBitStack2(t *testing.T) {
	testBoundedStack(t, "bitStack2a", BitStack2a_Empty, 2)
	testBoundedStack(t, "bitStack2b", BitStack2b_Empty, 2)
	testBoundedStack(t, "bitStack2c", BitStack2c_Empty, 2)
	testBoundedStack(t, "bitStack2d", BitStack2d_Empty, 63)
}
