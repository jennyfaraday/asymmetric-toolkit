package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestCliParser_HelpShort(t *testing.T) {
	var cfg Configuration
	args := []string{"-h"}
	if !cfg.Parse(args) {
		t.Errorf("Expected true on -h")
	}
}

func TestCliParser_HelpLong(t *testing.T) {
	var cfg Configuration
	args := []string{"--help"}
	if !cfg.Parse(args) {
		t.Errorf("Expected true on --help")
	}
}

func TestCliParser_Version(t *testing.T) {
	var cfg Configuration
	args := []string{"--version"}
	if !cfg.Parse(args) {
		t.Errorf("Expected true on --version")
	}
}

func TestCliParser_DomainEmpty(t *testing.T) {
	var cfg Configuration
	args := []string{"--domain"}
	if !cfg.Parse(args) {
		t.Errorf("Expected error since only the domain was given.")
	}
}

func TestCliParser_NoArgs(t *testing.T) {
	var cfg Configuration
	var args []string
	if !cfg.Parse(args) {
		t.Errorf("Expected an error")
	}
}

func TestCliParser_DomainOnly(t *testing.T) {
	var cfg Configuration
	args := []string{"--domain", "google.com"}
	if !cfg.Parse(args) {
		t.Errorf("Expected an error")
	}
}
func TestCliParser_DomainMode(t *testing.T) {
	var cfg Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence"}
	defer func(){recover()}()
	_ = cfg.Parse(args)
	t.Errorf("Expected an error")
}

func TestCliParser_DomainSequence(t *testing.T) {
	var cfg Configuration
	domainStr := "google.com"
	dnsServer := "udp:127.0.0.1:53"
	args := []string{"--domain", domainStr, "--mode", "sequence", "--dnsServer", dnsServer}
	if cfg.Parse(args) {
		//We hit an unexpected terminate result.
		t.Errorf("Error parsing sequence with no optionals. Args: %v", args)
	} else {
		errors.Assert(cfg.Domain.Get() == domainStr, fmt.Sprintf("Expected domain string not found. domain:'%s' expected:'%s'", domainStr, cfg.Domain.Get()))
		errors.Assert(cfg.Mode.IsSequence(), "Expected Sequence mode")
		errors.Assert(!cfg.Debug, "Expected debug to be false")
		errors.Assert(!cfg.Force, "Expected force to be false.")
		errors.Assert(cfg.Concurrency == defaultConcurrency, "Expected defaultConcurrency")
		errors.Assert(cfg.Delay == 0, "Expected delay not found.")
		errors.Assert(cfg.Depth == defaultDepth, "Expected depth not found.")
		errors.Assert(cfg.Dictionary == "", "Expected empty dictionary path/filename")
		errors.Assert(cfg.TargetServer.String() == dnsServer, "Expected dns servers not found.")
		errors.Assert(cfg.Output == "", "Expected empty output filename.")
		errors.Assert(cfg.Pattern.String() == defaultFilterPattern, "Expected filter pattern not found.")
		errors.Assert(cfg.RecordTypes.String() == defaultDnsRecordTypes, "Expected dns record types not found.")
		errors.Assert(cfg.Timeout == defaultTimeout, "Expected timeout not found.")
		errors.Assert(cfg.WordSize == defaultWordSize, "Expected wordsize not found.")
	}
}

func TestCliParser_DomainRandom(t *testing.T) {
	var cfg Configuration
	domainStr := "google.com"
	dnsServer := "udp:127.0.0.1:53"
	args := []string{"--domain", domainStr, "--mode", "sequence", "--dnsServer", dnsServer}
	if cfg.Parse(args) {
		//We hit an unexpected terminate result.
		t.Errorf("Error parsing random with no optionals")
	} else {
		errors.Assert(cfg.Domain.Get() == domainStr, fmt.Sprintf("Expected domain string not found. domain:'%s' expected:'%s'", domainStr, cfg.Domain.Get()))
		errors.Assert(cfg.Mode.IsSequence(), "Expected Sequence mode")
		errors.Assert(!cfg.Debug, "Expected debug to be false")
		errors.Assert(!cfg.Force, "Expected force to be false.")
		errors.Assert(cfg.Concurrency == defaultConcurrency, "Expected defaultConcurrency")
		errors.Assert(cfg.Delay == 0, "Expected delay not found.")
		errors.Assert(cfg.Depth == defaultDepth, "Expected depth not found.")
		errors.Assert(cfg.Dictionary == "", "Expected empty dictionary path/filename")
		errors.Assert(cfg.TargetServer.String() == dnsServer, "Expected dns servers not found.")
		errors.Assert(cfg.Output == "", "Expected empty output filename.")
		errors.Assert(cfg.Pattern.String() == defaultFilterPattern, "Expected filter pattern not found.")
		errors.Assert(cfg.RecordTypes.String() == defaultDnsRecordTypes, "Expected dns record types not found.")
		errors.Assert(cfg.Timeout == defaultTimeout, "Expected timeout not found.")
		errors.Assert(cfg.WordSize == defaultWordSize, "Expected wordsize not found.")
	}
}