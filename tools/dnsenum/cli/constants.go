package cli

import "asymmetric-effort/asymmetric-toolkit/tools/common/types"

const (
	Version               string                = "0.0.1"
	defaultConcurrency    types.PositiveInteger = 1  //default number of concurrent queries to run.
	defaultDepth          types.PositiveInteger = 1  //default number of DNS subdomain levels to attack
	maxDepth              types.PositiveInteger = 20 //Maximum number of DNS subdomain levels to attack
	defaultTimeout        types.PositiveInteger = 60 //Default number of seconds before connection timeout.
	defaultFilterPattern  string                = `.+`
	defaultDnsRecordTypes string                = "A,AAAA,MX,CNAME,NS,TXT,SOA,SRV"
	defaultWordSize       types.PositiveInteger = 5 //For a given sequence or random string, this is the max length of the 'word'
	DnsChars              string                = "WMEFSABCDGHIJKLNOPQRTUVXYZwmefsabcdghijklnopqrtuvxyz0123456789._-"
	SourceBufferSz        int                   = 1048576 //Size of the source buffer which feeds the payload (attack) function.
	ResponseBufferSz      int                   = 1048576 //Size of the response buffer for responses pending report processing.

	expectFlag  int = 0
	expectValue int = 1

	ExitTerminate bool = true
	ExitParseOk   bool = false
)
