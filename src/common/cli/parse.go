package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"fmt"
	"regexp"
)

func (o *Configuration) Parse(cliArguments []string) bool {
	o.LoadDefault()
	expected := ExpectFlag
	lastFlag := NoFlag
	for _, args := range cliArguments {
		switch expected {
		case ExpectFlag:
			expected = ExpectValue
			switch args {
			case "-h", "--help":
				lastFlag = UsageFlag
			case "--version", "-v":
				lastFlag = VersionFlag
			case "--concurrency":
				lastFlag = ConcurrencyFlag
			case "--debug":
				lastFlag = DebugFlag
				expected = ExpectFlag
				o.Debug = true
			case "--delay":
				lastFlag = DelayFlag
			case "--depth":
				lastFlag = DepthFlag
			case "--dictionary":
				lastFlag = DictionaryFlag
			case "--dnsServer":
				lastFlag = TargetServerFlag
			case "--domain":
				lastFlag = DomainFlag
			case "--force":
				lastFlag = ForceFlag
				expected = ExpectFlag
				o.Force = true
			case "--maxWordCount":
				lastFlag = MaxWordCountFlag
			case "--mode":
				lastFlag = ModeFlag
			case "--output":
				lastFlag = OutputFlag
			case "--pattern":
				lastFlag = PatternFlag
			case "--recordTypes":
				lastFlag = RecordTypesFlag
			case "--timeout":
				lastFlag = TimeoutFlag
			case "--wordSize":
				lastFlag = WordSizeFlag
			default:
				errors.Fatal(1, fmt.Sprintf("Encountered unexpected argument: %s", args))
			}
		case ExpectValue:
			func() {
				re := regexp.MustCompile(`^--.+$`)
				if re.MatchString(args) {
					errors.Fatal(1, "Expected value, not flag.")
				}
			}()
			switch lastFlag {
			case NoFlag, UsageFlag:
				ShowUsage()
				return ExitTerminate
			case VersionFlag:
				ShowVersion()
				return ExitTerminate
			case ConcurrencyFlag:
				o.Concurrency.Set(args)
				expected = ExpectFlag
			case DebugFlag:
				expected = ExpectFlag
			case DepthFlag:
				expected = ExpectFlag
				o.Depth.Set(args)
			case DelayFlag:
				expected = ExpectFlag
				o.Delay.Set(args)
			case DictionaryFlag:
				expected = ExpectFlag
				o.Dictionary.Set(args)
			case DomainFlag:
				expected = ExpectFlag
				o.Domain.Set(args)
			case TargetServerFlag:
				expected = ExpectFlag
				o.TargetServer.Set(args)
			case ForceFlag:
				expected = ExpectFlag
			case MaxWordCountFlag:
				o.MaxWordCount.Set(args)
				expected = ExpectFlag
			case ModeFlag:
				o.Mode.Set(args)
				expected = ExpectFlag
			case OutputFlag:
				o.Output.Set(args)
				expected = ExpectFlag
			case PatternFlag:
				o.Pattern.Set(args)
				expected = ExpectFlag
			case RecordTypesFlag:
				o.RecordTypes.Set(args)
				expected = ExpectFlag
			case TimeoutFlag:
				o.Timeout.Set(args)
				expected = ExpectFlag
			case WordSizeFlag:
				o.WordSize.Set(args)
				expected = ExpectFlag
			default:
				panic("invalid flag")
			}
		default:
			panic(
				fmt.Sprintf("Expected either expectFlag or expectValue.  Encountered: %v",
					expected))
		}
	} /* end for */
	//
	//Perform a final validation...
	//
	if lastFlag == UsageFlag {
		ShowUsage()
		return ExitTerminate
	}
	if lastFlag == VersionFlag {
		ShowVersion()
		return ExitTerminate
	}
	if o.Domain.Get() == "" {
		fmt.Println("Missing domain (required).  Use --domain <string> to specify.")
		return ExitTerminate
	}

	if o.Mode.Get() == sourcetype.NotSet {
		fmt.Println("Missing mode (required).  Use --mode <sequence|random|dictionary> to specify.")
		return ExitTerminate
	}

	if o.TargetServer.Get() == "" {
		fmt.Println("Missing dnsServer (required).  Use --dnsServer <udp|tcp>:<ipaddr>:<port> to specify.")
		return ExitTerminate
	}

	if o.Mode.IsDictionary() {
		if o.Dictionary == "" {
			fmt.Println("Missing Dictionary.  Use --dictionary <string> when --mode dictionary is used.")
			return ExitTerminate
		}
		if !o.Dictionary.Exists() {
			fmt.Printf("Dictionary file not found (%s)\n", o.Dictionary)
			return ExitTerminate
		}
		if o.WordSize >= 0 {
			fmt.Printf("In Dictionary mode --wordSize is not allowed")
			return ExitTerminate
		}
		if o.MaxWordCount >= 0 {
			fmt.Printf("In Dictionary mode --maxWordCount is not allowed")
			return ExitTerminate
		}
	}
	if o.Mode.IsRandom() || o.Mode.IsSequence() {
		if o.Dictionary != "" {
			fmt.Println("Do not use --dictionary <file> with random or sequential mode.")
			return ExitTerminate
		}
		if o.MaxWordCount == 0 {
			o.MaxWordCount = types.PositiveInteger(int(o.WordSize) * len(DNSChars))
		}
		if o.WordSize == 0 {
			o.WordSize = DefaultWordSize
		}
	}
	if o.Output != "" {
		if !o.Force && o.Output.Exists() {
			fmt.Println("Output file exists.  Use --force to overwrite")
			return ExitTerminate
		}
	}
	if o.Depth > MaxDepth {
		fmt.Printf("Depth (--depth) exceeds maxDepth (%d)\n", MaxDepth)
		return ExitTerminate
	}
	return ExitParseOk
}