package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestCommandLine_ParseHelp(t *testing.T) {
	//happy path --help
	args := []string{"--help"}
	var spec Specification = Specification{
		Author:      "Sam Caldwell",
		AuthorEmail: "mail@samcaldwell.net",
		ProgramName: "Asymmetric Effort",
		Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		Version:     "0.0.1",
		Argument: map[string]ArgumentDescriptor{
			"flag": {
				1000, // >= 1000 is a project-defined FlagId
				None,
				"false",
				"This is a flag.",
				ParserFlag("flag"),
				ExpectNone, //We expect None (no value) which will expect a flag in the end.
			},
			"intVal": {
				1001, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This is an integer value",
				ParserInt(),
				ExpectValue,
			},
			"intValDefault": {
				1002, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This tests default value (5)",
				ParserInt(0, 10),
				ExpectValue,
			},
			"strList": {
				1003, // >= 1000 is a project-defined FlagId
				List,
				"A,B,C,D",
				"This tests a default list",
				ParserList(","),
				ExpectValue,
			},
			"strVal": {
				1004, // >= 1000 is a project-defined FlagId
				String,
				"defaultValueShouldNotBeThere",
				"This is a string value.",
				ParserString(),
				ExpectValue,
			},
		},
	}
	spec.AddHelp()
	var ui CommandLine
	_, err := ui.Parse(&spec, &args)
	if err != nil {
		panic(err)
	}
}

func TestCommandLine_ParseVersion(t *testing.T) {
	//happy path --version
	args := []string{"--version"}
	var spec Specification = Specification{
		Author:      "Sam Caldwell",
		AuthorEmail: "mail@samcaldwell.net",
		ProgramName: "Asymmetric Effort",
		Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		Version:     "0.0.1",
		Argument: map[string]ArgumentDescriptor{
			"flag": {
				1000, // >= 1000 is a project-defined FlagId
				None,
				"false",
				"This is a flag.",
				ParserFlag("flag"),
				ExpectNone, //We expect None (no value) which will expect a flag in the end.
			},
			"intVal": {
				1001, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This is an integer value",
				ParserInt(),
				ExpectValue,
			},
			"intValDefault": {
				1002, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This tests default value (5)",
				ParserInt(0, 10),
				ExpectValue,
			},
			"strList": {
				1003, // >= 1000 is a project-defined FlagId
				List,
				"A,B,C,D",
				"This tests a default list",
				ParserList(","),
				ExpectValue,
			},
			"strVal": {
				1004, // >= 1000 is a project-defined FlagId
				String,
				"defaultValueShouldNotBeThere",
				"This is a string value.",
				ParserString(),
				ExpectValue,
			},
		},
	}
	spec.AddVersion()
	var ui CommandLine
	_, err := ui.Parse(&spec, &args)
	if err != nil {
		panic(err)
	}
}

func TestCommandLine_ParseCustom(t *testing.T) {
	args := []string{"--flag", "--intVal", "10", "--strVal", "myValue"}
	var spec Specification = Specification{
		Author:      "Sam Caldwell",
		AuthorEmail: "mail@samcaldwell.net",
		ProgramName: "Asymmetric Effort",
		Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		Version:     "0.0.1",
		Argument: map[string]ArgumentDescriptor{
			"flag": {
				1000, // >= 1000 is a project-defined FlagId
				None,
				"false",
				"This is a flag.",
				ParserFlag("flag"),
				ExpectNone, //We expect None (no value) which will expect a flag in the end.
			},
			"intVal": {
				1001, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This is an integer value",
				ParserInt(),
				ExpectValue,
			},
			"intValDefault": {
				1002, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This tests default value (5)",
				ParserInt(0, 10),
				ExpectValue,
			},
			"strList": {
				1003, // >= 1000 is a project-defined FlagId
				List,
				"A,B,C,D",
				"This tests a default list",
				ParserList(","),
				ExpectValue,
			},
			"strVal": {
				1004, // >= 1000 is a project-defined FlagId
				String,
				"defaultValueShouldNotBeThere",
				"This is a string value.",
				ParserString(),
				ExpectValue,
			},
		},
	}
	var ui CommandLine
	_, err := ui.Parse(&spec, &args)
	if err != nil {
		panic(err)
	}

	errors.Assert(ui.Arguments[FlagHelp] == nil, "Expected nil help argument")

	errors.Assert(ui.Arguments[FlagVersion] == nil, "Expected nil version argument")

	errors.Assert(ui.Arguments[FlagDebug] != nil, "Expected non-nil debug argument")

	errors.Assert(!ui.Arguments[FlagDebug].Boolean(), "Expected false debug argument")
	/*
		errors.Assert(ui.Arguments[FlagForce] == nil , "Expected nil force argument")
	*/
	errors.Assert(
		ui.Arguments[1000].Type == Boolean,
		fmt.Sprintf("1.Expected type(%v): %v", Boolean, ui.Arguments[1000].Type))

	errors.Assert(
		ui.Arguments[1001].Type == Integer,
		fmt.Sprintf("2.Expected type(%v): %v", Integer, ui.Arguments[1001].Type))

	errors.Assert(
		ui.Arguments[1002].Type == Integer,
		fmt.Sprintf("3.Expected type(%v): %v", Integer, ui.Arguments[1002].Type))

	errors.Assert(
		ui.Arguments[1000].String() == "false",
		fmt.Sprintf("Unexpected Value:%v", ui.Arguments[1000].String()))

	errors.Assert(
		ui.Arguments[1001].String() == "10",
		fmt.Sprintf("Unexpected Value:%v", ui.Arguments[1001].String()))

	errors.Assert(
		ui.Arguments[1002].String() == "5",
		fmt.Sprintf("Unexpected Value:%v", ui.Arguments[1002].String()))

	errors.Assert(
		ui.Arguments[1003].String() == "A,B,C,D",
		fmt.Sprintf("Unexpected Value:%v", ui.Arguments[1003].String()))

	errors.Assert(
		ui.Arguments[1004].Type == String,
		fmt.Sprintf("4.Expected type(%v): %v", String, ui.Arguments[1004].Type))

	errors.Assert(
		ui.Arguments[1004].String() == "myValue",
		fmt.Sprintf("Unexpected Value:%v", ui.Arguments[1004].String()))
}

func TestCommandLine_ParseConcurrency(t *testing.T) {
	//happy path --concurrency <int>
	args := []string{"--concurrency", "10"}
	var spec Specification = Specification{
		Author:      "Sam Caldwell",
		AuthorEmail: "mail@samcaldwell.net",
		ProgramName: "Asymmetric Effort",
		Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		Version:     "0.0.1",
		Argument:    map[string]ArgumentDescriptor{},
	}
	spec.AddConcurrency(20)
	var ui CommandLine
	_, err := ui.Parse(&spec, &args)
	if err != nil {
		panic(err)
	}
	type Configuration struct {
		Concurrency int
	}
	var config Configuration
	config.Concurrency = ui.Arguments[FlagConcurrency].Integer()
}
