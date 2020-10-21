package LogFacility_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"testing"
)

func TestLogFacilityGet_Happy(t *testing.T){
	var f LogFacility.Facility
	f= LogFacility.Facility("MyLog") //Happy
	errors.Assert(string(f)=="MyLog", "Expected value not properly stored.")
}

func TestLogFacilityGet_Bad(t *testing.T){
	var f LogFacility.Facility
	defer func(){recover()}()
	f= LogFacility.Facility("MyBadLog!")
	_ = f.Get()
	t.Errorf("Expected error")
}

func TestLogFacilityGet_Empty(t *testing.T){
	var f LogFacility.Facility
	defer func(){recover()}()
	f= LogFacility.Facility("MyBadLog!")
	_ = f.Get()
	t.Errorf("Expected error")
}
