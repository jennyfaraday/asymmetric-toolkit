package cli

import "strconv"

/*
	Specification::AddConcurrency() implements --concurrency <int> flags.
*/
const (
	concurrencyHelpText = "Indicates the concurrency (number of attackers) to be allowed."
	concurrencyArgLong  = "concurrency"
)

func (o *Specification) AddConcurrency(defaultValue int) {
	//
	// Initialize the Argument object.
	//
	if defaultValue == 0 {
		panic("Concurrency cannot be zero (0)")
	}
	o.Initialize()
	o.Argument[concurrencyArgLong] = ArgumentDescriptor{
		FlagConcurrency,
		Integer,
		strconv.Itoa(defaultValue),
		concurrencyHelpText,
		ParserInt(0,65536),
		ExpectValue,
	}
}
