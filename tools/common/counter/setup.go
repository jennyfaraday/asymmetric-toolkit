package counter

import "asymmetric-effort/asymmetric-toolkit/tools/common/errors"

func (o *Counter) Setup(charset string, wordSize int) {
	errors.Assert(charset != "", "Counter::Setup() expects non empty string for charset")
	errors.Assert(wordSize > 0, "Counter::Setup() expects positive integer word size.")
	o.runes = func() *[]rune { d := []rune(charset); return &d }()
	o.maxBase = uint8(len(*o.runes))
	o.data = func() *[]uint8 { d := make([]uint8, wordSize); return &d }()
}