package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"fmt"
	"testing"
)

func TestSourceGenerateSequence(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "0123456789"
	const maxWordSize = 9
	var s source.Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence", "--dnsServer", "udp:127.0.0.1:53"}
	config.Parse(args)
	s.Config = &config

	for s.Config.WordSize = 1; s.Config.WordSize < maxWordSize; s.Config.WordSize++ {
		s.Config.MaxWordCount = 1000000000
		s.AllowedChars = func() *string { str := keyspace; return &str }()
		s.Feed.Setup(cli.SourceBufferSz * 1000)
		/*
			Run Generator
		*/
		fmt.Printf("Starting generator.  Size:%d of %d\n", s.Config.WordSize, maxWordSize)
		s.GenerateSequence()
		s.Feed.Close()
		fmt.Printf("Finished generating.  Number items in channel: %d\n", s.Feed.Length())
		/*
			Analyze Result
		*/
		expectedCount := s.Feed.Length()
		last := ""
		for i := 0; i <= expectedCount; i++ {
			if s.Feed.Length() > 0 {
				last = s.Feed.Pop()
			}
		}
		errors.Assert(s.Feed.Length() == 0, "Expected to have consumed all elements")
		fmt.Printf("Consumed queue of %d elements (last:%s)\n", expectedCount, last)
	}
	fmt.Println("TestSourceGenerateSequence...Done")
}

func TestSourceGenerateSequenceBinary(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "01"
	const maxWordSize = 9
	var s source.Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence", "--dnsServer","udp:127.0.0.1:53"}
	config.Parse(args)
	s.Config = &config
	for s.Config.WordSize = 1; s.Config.WordSize < maxWordSize; s.Config.WordSize++ {
		s.Config.MaxWordCount = 1000000000
		s.AllowedChars = func() *string { str := keyspace; return &str }()
		s.Feed.Setup(cli.SourceBufferSz * 1000)
		/*
			Run Generator
		*/
		fmt.Printf("Starting generator.  Size:%d of %d\n", s.Config.WordSize, maxWordSize)
		s.GenerateSequence()
		s.Feed.Close()
		fmt.Printf("Finished generating.  Number items in channel: %d\n", s.Feed.Length())
		/*
			Analyze Result
		*/
		expectedCount := s.Feed.Length()
		last := ""
		for i := 0; i <= expectedCount; i++ {
			if s.Feed.Length() > 0 {
				last = s.Feed.Pop()
			}
		}
		errors.Assert(s.Feed.Length() == 0, "Expected to have consumed all elements")
		fmt.Printf("Consumed queue of %d elements (last:%s)\n", expectedCount, last)
	}
	fmt.Println("TestSourceGenerateSequenceBinary...Done")
}
