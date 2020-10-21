package misc_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"fmt"
	"strconv"
	"testing"
)

func TestMiscIsPort(t *testing.T){
	for i:=0;i<65536;i++{
		errors.Assert(misc.IsPort(strconv.Itoa(i)),fmt.Sprintf("Expected port (%d) to be valid",i))
	}
	errors.Assert(!misc.IsPort("-1"),fmt.Sprintf("Expected port >= 0"))
	errors.Assert(!misc.IsPort("65536"),fmt.Sprintf("Expected port <=65535"))
}