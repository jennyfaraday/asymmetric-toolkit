package LogFacility

import (
	"regexp"
	"testing"
)

func TestLogFacilityConstantsFacRegEx(t *testing.T){
	_ = regexp.MustCompile(facilityRegExPattern)
}