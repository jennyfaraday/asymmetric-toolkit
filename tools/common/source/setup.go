package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
)

func (o *Source) Setup(config *cli.Configuration, bufferSz int, allowedChars string) {
	errors.Assert(config != nil, "Encountered nil configuration in Source::Setup()")
	o.config = config

	errors.Assert(bufferSz > 1, "Expected sourceBufferSz > 1")
	o.feed = make(chan string, bufferSz)

	errors.Assert(allowedChars != "", "Expected non-empty string in allowedChars")
	o.allowedChars = &allowedChars

	o.isPaused = true //By default the feed will be paused until the owner unpauses it.

	switch {
	/*
		Startup the relevant generator function (paused by the isPaused flag).
	*/
	case config.Mode.IsSequence():
		go o.generateSequence()
	case config.Mode.IsRandom():
		go o.generateRandom()
	case config.Mode.IsDictionary():
		go o.generateDictionary()
	default:
		panic("Source::Setup() encountered Mode NotSet")
	}
}
