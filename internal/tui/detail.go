package tui

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"time"
)

type detailModel struct {
	sub     chan tea.Msg // where we'll receive activity notifications
	content bytes.Buffer
	i       item

	stop    chan struct{}
	spinner spinner.Model
	back    tea.Model
}

func newDetailModel(back tea.Model, i item) *detailModel {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := &detailModel{
		sub:     make(chan tea.Msg, 1),
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
		m.content.WriteString(v)         // record external activity
		return m, waitForActivity(m.sub) // wait for next event
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case error:
		if !errors.Is(v, io.EOF) {
			m.content.WriteString(v.Error())
		}
		// TODO: render spinner stopped.
		return m, nil
	default:
		return m, nil
	}
}

func (m *detailModel) View() string {
	s := fmt.Sprintf(`
Analyze Slow Query:
%s
%s Received:
%s

Press any key to return`,
		m.i.Query(), m.spinner.View(), m.content.String())
	return s
}

func (m *detailModel) listenForHandler(handler Handler) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		recv := handler(ctx, m.i.LogMessage)
		for {
			select {
			case <-m.stop:
				close(m.sub)
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
