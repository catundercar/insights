package analyzer

import (
	"context"
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ollama/ollama/api"
	"insights"
	"insights/pkg/prompt"
	"io"
	"net/http"
	"net/url"
)

type Config struct {
	ApiServer string
}

type Analyzer struct {
	client *api.Client
}

func NewAnalyzer(c Config) (*Analyzer, error) {
	schema, err := url.Parse(c.ApiServer)
	if err != nil {
		return nil, err
	}

	analyzer := new(Analyzer)
	analyzer.client = api.NewClient(schema, http.DefaultClient)
	return analyzer, nil
}

func (a *Analyzer) HandleCompletion(ctx context.Context, log insights.LogMessage) <-chan tea.Msg {
	ch := make(chan tea.Msg, 1)
	go func() {
		defer close(ch)
		pt, err := prompt.GenPrompt(prompt.NewPromptData(log.DatabaseType(), log.Message()))
		if err != nil {
			ch <- err
			return
		}
		req := api.GenerateRequest{
			Model:  "qwen2",
			Prompt: pt,
		}
		err = a.client.Generate(ctx, &req, func(resp api.GenerateResponse) error {
			if !resp.Done {
				ch <- resp.Response
				return nil
			}
			return io.EOF
		})
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				ch <- err
			}
		}
	}()
	return ch
}
