package cli

/*
	The ArgumentType indicates the expected datatype of the commandline argument once all parsing and pre-processing
	is done.  This is used to validate the input and allow for its extraction to the final state before the command
	line processor terminates and its memory is free.

	The ArgumentType is limited to the primitive types.
*/
type ArgumentType int

const (
	None          ArgumentType = 0
	String        ArgumentType = 1
	Integer       ArgumentType = 2
	Float         ArgumentType = 3
	Boolean       ArgumentType = 4
	List          ArgumentType = 5

	argumentTypes              = "None,String,Integer,Float,Boolean,List"

	comma                      = ","
)
