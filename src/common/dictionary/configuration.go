package dictionary

type Configuration struct {
	FileName       string // Descriptor path / filename
	Overwrite      bool   // Flag to indicate if any existing file is to be overwritten.
	FormatVersion  Version
	ScoreVersion   Version
	Passphrase     []byte
	CompressionAlg ioCompression
}