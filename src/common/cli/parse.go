package cli

import (
	"fmt"
	"os"
)

func (o *CommandLine) Parse(spec *Specification) (exit bool, err error) {
	/*
		Parse the existing command line arguments, perform validation and store the
		values into the internal state.

		Input: *Specification (describe how the argument parser will work.
		Output:
			- Should the caller exit or continue? (default: continue.
	*/
	var expected NextExpected = ExpectFlag
	var lastFlag *ArgumentDescriptor = nil

	spec.AddUsage()
	spec.AddVersion()

	for _, arg := range os.Args[1:] {
		//
		// Iterate through all arguments on the commandline.  We expect either
		// a long flag (--flag), a short flag (-f) or a contiguous string value.
		//
		switch expected {
		case ExpectFlag:
			//
			// If we expect a flag (--long or -l), we will parse the flag and
			// determine its specifications to process the next argument.
			//
			if knownSpec, ok := spec.Argument[stripPrefix(arg)]; ok {
				//
				// For the current argument, does it exist in our specification's argument map?
				// If it does exist, we have a detected flag we can process...
				//
				lastFlag = &knownSpec        // Save the current ArgumentDescriptor for the flag.
				expected = knownSpec.Expects // Route the next argument as expected in the spec.
				if expected == ExpectNone {
					// If we expect nothing (e.g. --help, --debug, --verbose) where the flag is the value
					// in and of itself, we simply process it now rather than wait another cycle (which would process
					// the next argument).
					if err, o.Arguments[lastFlag.FlagId] = lastFlag.Parse(); err == nil {
						if o.Arguments[lastFlag.FlagId] == nil {
							// If our flag Parse() function returned nil in this case
							// we would know the Parse() function is terminal and we should exit without error,
							// as in the --help or --version use cases.
							return true, nil
						}
						expected = ExpectFlag // Reset and get another flag.
					} else {
						return true, err
					}
				}
				continue
			} else {
				//
				// Terminate with an error.  We failed to find the current argument as an
				// argument in our current specification.
				//
				return true, fmt.Errorf("missing or unrecognized argument: %s", arg)
			}
			//
		case ExpectValue:
			//
			// If we expect a value, we presume a flag was detected first to which we can map the
			// parsed value currently in arg.
			//
			if lastFlag == nil {
				//
				// If we have no argument (arg) flag, we return an error because we cannot map a value
				// without a flag to represent what the value actually means.
				//
				return true, fmt.Errorf("a value was found with no flag: %s", arg)
			}
			//
			// Our lastFlag (ArgumentDescriptor) defines how we should proceed to parse, validate and
			// handle the value currently in arg.  If successful, the CommandLine (o) structure will
			// contain the parsed (string) value in the Argument map.
			//
			if lastFlag.Parse == nil {
				//
				// We have a nil pointer reference for our parser.  To avoid a crash we simply
				// continue at this point and assume a noop.  No value will be stored for this
				// argument and move on to the next flag.
				//
				expected = ExpectFlag
				continue
			} else {
				var err error
				//
				// If a Parse() function is provided for validation, we will execute the same
				// to parse and validate our data specific to its context. If no error occurs,
				// We expect the Parse() function will return an Argument object containing the
				// processed value and any additional features.
				//
				if err, o.Arguments[lastFlag.FlagId] = lastFlag.Parse(); err == nil {
					expected = ExpectFlag
				} else {
					//
					// For each ArgumentDescriptor, a Parse() function is provided.  This function
					// will perform any context-specific validation and store the value as an Argument
					// object in CommandLine (o).
					//
					return true, err
				}
				if o.Arguments[lastFlag.FlagId] == nil {
					//
					// Perform a nil pointer check.  Parse() functions cannot return nil Arguments
					// without returning an error.  A nil Argument implies that the argument failed
					// to parse/validate.
					//
					return true, fmt.Errorf("our flag parse() function cannot return a nil "+
						"Argument pointer (%d)", lastFlag.FlagId)
				}
			}
		case ExpectEnd:
			//No other arguments will be processed.  We expect to exit.
			return true, nil
		default:
			return true, fmt.Errorf("unexpected argument parser state. Expected: %d", expected)
		}
	}
	return false, nil
}

func stripPrefix(arg string) string {
	//ToDo: Strip prefix
	return arg
}
