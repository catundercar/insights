package insights

import (
	"context"
	insights_v1 "insights/api/insights/v1"
	"time"
)

type collectorOption struct {
	Duration time.Duration
}

type Collector interface {
	Collect(ctx context.Context) ([]*insights_v1.SlowQueryLog, error)
}
