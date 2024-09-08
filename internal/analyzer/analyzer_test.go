package analyzer

import (
	"context"
	"fmt"
	"insights"
	"insights/internal/tui"
	"testing"
)

type message struct {
	query     string
	frequency int
}

func (m message) DatabaseType() string {
	return "MySQL"
}

func (m message) Query() string {
	return m.query
}

func (m message) Message() string {
	return "message: " + m.query
}

func (m message) Frequency() int {
	return m.frequency
}

func TestNewAnalyzer(t *testing.T) {
	messages := []insights.LogMessage{
		message{
			query:     "SELECT count(*) FROM users",
			frequency: 4,
		},
		message{
			query:     "SELECT * FROM users",
			frequency: 1,
		},
		message{
			query:     "SELECT * FROM users WHERE id = 1",
			frequency: 6,
		},
		message{
			query:     "SELECT * FROM users WHERE id = 1 AND name = 'John'",
			frequency: 3,
		},
	}

	analyzer, err := NewAnalyzer(Config{ApiServer: "http://127.0.0.1:11434"})
	if err != nil {
		t.Fatal(err)
	}
	m := tui.NewModel(analyzer.HandleCompletion)
	m.Render(messages)
}

func TestAnalyzer_HandleCompletion(t *testing.T) {
	analyzer, err := NewAnalyzer(Config{ApiServer: "http://127.0.0.1:11434"})
	if err != nil {
		t.Fatal(err)
	}
	ch := analyzer.HandleCompletion(context.TODO(), message{query: "SELECT * FROM users WHERE id = 1"})
	for v := range ch {
		fmt.Print(v)
	}
}
