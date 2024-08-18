package mongo

const (
	SlowQueryMsg = "Slow query"
)

// LogMessage is the message of [mongodb](https://www.mongodb.com/docs/manual/reference/log-messages/#log-message-field-types).
type LogMessage struct {
	//Timestamp time.Time `json:"t" bson:"t"`       // The Timestamp field type indicates the precise date and time at which the logged event occurred.
	//Severity  string    `json:"s" bson:"s"`       // The Severity field type indicates the severity level associated with the logged event.
	//Component string    `json:"c" bson:"c"`       // TODO
	//Attr      bson.D    `json:"attr" bson:"attr"` // TODO
	Msg string `json:"msg" bson:"msg"`
}
