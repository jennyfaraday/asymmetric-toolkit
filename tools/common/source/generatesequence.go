package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/counter"
)

func (o *Source) generateSequence() {
	/*
		Generate words as a sequence within a given keyspace (allowedChars).
		For example: A B C AA AB AC BA BB BC CA CB CC and so on where the
		length of the words is determined by config.WordSize.
	 */
	checkWordSize := func(wsz int) bool {
		return wsz <= int(o.config.WordSize)
	}
	checkWordCount := func() bool {
		return o.counter <= int(o.config.MaxWordCount)
	}
	for wordSize := 1; checkWordSize(wordSize) && checkWordCount(); wordSize++ {
		var generator counter.Counter
		generator.Setup(*o.allowedChars, wordSize)
		for generator.Increment(0) && checkWordCount() { //Iterate through the field of wordSize chars.
			o.feed <- generator.String() //Produce a new word for each one and push to the feed.
			o.counter++
		}
	}
}
