package tui

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"insights"
	"testing"
	"time"
)

type message struct {
	query     string
	frequency int
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

func TestRender(t *testing.T) {
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
	m := NewModel(func(ctx context.Context, log insights.LogMessage) <-chan tea.Msg {
		ch := make(chan tea.Msg, 1)
		ch <- fmt.Sprintf("log: %s, Frequency: %d", log.Query(), log.Frequency())
		go func() {
			tk := time.NewTicker(300 * time.Millisecond)
			for {
				select {
				case <-tk.C:
					ch <- "test "
				case <-ctx.Done():
					tk.Stop()
					close(ch)
					return
				}
			}
		}()
		return ch
	})
	m.Render(messages)
}
