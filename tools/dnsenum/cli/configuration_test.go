package cli

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestConfigurationStruct(t *testing.T){
	var c Configuration
	errors.Assert(c.Concurrency==0,"Expect 0")
	errors.Assert(!c.Debug, "expect false")
	errors.Assert(c.Delay == 0, "expect 0")
	errors.Assert(c.Depth == 0, "expect 0")
	errors.Assert(c.Dictionary == "", "expect empty string")
	errors.Assert(c.TargetServer == "", "expect empty list")
	errors.Assert(c.Domain == "", "expect empty string")
	errors.Assert(!c.Force, "Expect false")
	errors.Assert(c.Mode == 0, "expect 0")
	errors.Assert(c.Output == "", "expect empty string")
	errors.Assert(c.Pattern.String() == "", fmt.Sprintf("expect empty string.  Found: %s", c.Pattern.String()))
	errors.Assert(c.Pattern.String() == "", "expect empty string")
	errors.Assert(c.RecordTypes.String() == "", "expect empty list")
	errors.Assert(c.Timeout == 0, "expect 0")
	errors.Assert(c.WordSize == 0, "expect 0")
	errors.Assert(c.MaxWordCount == 0, "expect 0")
}