package casts

import "testing"

func TestBoolToInt(t *testing.T) {
	_ = BooltoInt(true)
	_ = BooltoInt(false)
}
