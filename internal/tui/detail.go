package tui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

type detailModel struct {
	sub     chan tea.Msg // where we'll receive activity notifications
	content string
	i       item

	stop    chan struct{}
	spinner spinner.Model
	back    tea.Model
}

func newDetailModel(sub chan tea.Msg, back tea.Model, i item) *detailModel {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := &detailModel{
		sub:     sub,
		i:       i,
		spinner: s,
		back:    back,
		stop:    make(chan struct{}),
	}
	return m
}

func (m *detailModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		waitForActivity(m.sub), // wait for activity
	)
}

func (m *detailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case tea.KeyMsg:
		close(m.stop)
		return m.back, nil
	case string:
		m.content += v                   // record external activity
		return m, waitForActivity(m.sub) // wait for next event
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m *detailModel) View() string {
	s := fmt.Sprintf("\n %s Analyze Slow Query: %s\n received: %s\n\n Press any key to return\n",
		m.spinner.View(), m.i.Query(), m.content)
	return s
}

func (m *detailModel) listenForHandler(handler Handler) tea.Cmd {
	return func() tea.Msg {
		// todo：request ollama.

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		recv := handler(ctx, m.i.LogMessage)
		for {
			select {
			case <-m.stop:
				return ""
			case msg := <-recv:
				m.sub <- msg
			}
		}
	}
}

func listenForActivity(stop chan struct{}, sub chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		// todo：request ollama.
		tk := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-stop:
				return ""
			case <-tk.C:
				sub <- "todo"
			}
		}
	}
}

// A command that waits for the activity on a channel.
func waitForActivity(sub chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		s := <-sub
		return s
	}
}
