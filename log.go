package insights

type LogMessage interface {
	Query() string
	Message() string
	Frequency() int
}

// RawLog is the raw data of logs.
type RawLog interface {
	GetRaw() []byte
}
