package types

type DataSource uint8

const (
	NotSet     DataSource = 0
	Sequence   DataSource = 1
	Random     DataSource = 2
	Dictionary DataSource = 3
	// See BadDataSource = 255 in dataSource_test.go
)

