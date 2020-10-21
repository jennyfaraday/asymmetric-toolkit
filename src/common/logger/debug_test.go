package logger_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestLoggerDebugWarn(t *testing.T) {
	var log logger.Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(level.Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Warningf("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+` +
		` [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(WARN\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}

func TestLoggerDebugCritical(t *testing.T) {
	var log logger.Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(level.Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Critical("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+` +
		` [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(CRIT\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}

func TestLoggerDebugError(t *testing.T) {
	var log logger.Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(level.Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Error("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+` +
		` [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(ERROR\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}


func TestLoggerDebugInfo(t *testing.T) {
	var log logger.Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(level.Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Infof("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+` +
		` [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(INFO\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}
