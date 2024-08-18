package mongo

import (
	"context"
	"insights"
	insights_v1 "insights/api/insights/v1"
)

type Collector struct {
	input insights.IteratorInput
	p     *parser
}

func NewCollector(input insights.IteratorInput) insights.Collector {
	return &Collector{
		input: input,
		p:     NewParser(),
	}
}
func (c *Collector) Collect(_ context.Context) ([]*insights_v1.SlowQueryLog, error) {
	slowLogs := make([]*insights_v1.SlowQueryLog, 0, 8)
	iter := c.input.Iter()
	for iter.Next() {
		raw := iter.Value()
		lm, err := c.p.ParseRaw(raw.GetRaw())
		if err != nil {
			return nil, err
		}
		if lm.Msg != SlowQueryMsg {
			continue
		}
		slowLogs = append(slowLogs, &insights_v1.SlowQueryLog{
			Content:  string(raw.GetRaw()),
			Database: insights_v1.SlowQueryLog_MONGODB,
		})
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return slowLogs, nil
}
