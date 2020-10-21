package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/entropy"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"fmt"
	"testing"
)

func TestSourceGenerateRandom(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var s source.Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "random", "--dnsServer", "udp:127.0.0.1:53", "--maxWordCount", "20"}
	config.Parse(args)
	s.Config = &config
	s.Config.WordSize = 100
	s.Config.MaxWordCount = 100
	s.AllowedChars = func() *string { str := keyspace; return &str }()
	s.Feed.Setup(cli.SourceBufferSz)
	/*
		Run Generator
	*/
	fmt.Println("Starting generator")
	s.GenerateRandom()
	s.Feed.Close()
	/*
		Analyze Result
	*/
	expectedCount := s.Feed.Length()
	word := ""
	totalShannons := 0
	minShannons := entropy.HighEntropyThreshold
	maxShannons := 0
	count := 0
	for i := 0; i <= expectedCount; i++ {
		if s.Feed.Length() > 0 {
			word = s.Feed.Pop()
			count++
			shannons := entropy.GetShannons(word)
			totalShannons += shannons
			if shannons < minShannons {
				minShannons = shannons
			}
			if shannons > maxShannons {
				maxShannons = shannons
			}
			fmt.Println("Shannons:", totalShannons, "shannons.  Avg", totalShannons/count)
		}
	}
	avgShannons := totalShannons / count
	errors.Assert(s.Feed.Length() == 0, "Expected to have consumed all elements")
	errors.Assert(float64(minShannons) > (0.90 * float64(entropy.HighEntropyThreshold)), fmt.Sprintf("Expected high min entropy (%d)",minShannons))
	errors.Assert(avgShannons > entropy.HighEntropyThreshold, "Expected high average entropy")
	errors.Assert(maxShannons > entropy.HighEntropyThreshold, "Expected high max entropy")
	errors.Assert(count == expectedCount, "Expected count match")
	fmt.Printf("Consumed queue of %d elements", expectedCount)
}
