package insights

type LogMessage interface {
	Query() string
	Message() string
	Frequency() int
	DatabaseType() string
}

// RawLog is the raw data of logs.
type RawLog interface {
	GetRaw() []byte
}
