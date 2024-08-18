package insights

// Input is the log data source from file or loki.
type Input interface {
	Fetch() ([]RawLog, error)
}

// IteratorInput wrappers Input like an iterator.
type IteratorInput interface {
	Iter() Iterator
}

type Iterator interface {
	Next() bool
	Value() RawLog
	Err() error
}
