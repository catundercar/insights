package tui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"insights"
	"os"
	"sort"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	insights.LogMessage
}

func (i item) Title() string { return i.Query() }

func (i item) Description() string {
	return i.Message() + fmt.Sprintf(" frequency: %d", i.Frequency())
}

func (i item) FilterValue() string { return i.Message() }

type Handler func(ctx context.Context, log insights.LogMessage) <-chan tea.Msg

type Model struct {
	sub     chan tea.Msg
	list    list.Model
	handler Handler
}

func NewModel(fn Handler) *Model {
	m := &Model{
		sub:     make(chan tea.Msg),
		handler: fn,
	}
	return m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				dm := newDetailModel(m.sub, m, i)
				return dm, tea.Batch(dm.Init(), dm.listenForHandler(m.handler))
			}
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return docStyle.Render(m.list.View())
}

func (m *Model) Render(messages []insights.LogMessage) {
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Frequency() > messages[j].Frequency()
	})

	items := make([]list.Item, 0, len(messages))
	for _, message := range messages {
		items = append(items, item{message})
	}

	m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.list.Title = "Slow Queries: "

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
