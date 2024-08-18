package prompt

import (
	"bytes"
	"fmt"
	"text/template"
)

type PromptData struct {
	// DatabaseName is the name of database, such as mysql, postgresql, etc.
	DatabaseName string
	// SlowQueryLog is the slow query log.
	SlowQueryLog string
	// Language is the language of result.
	Language string
}

func GenPrompt(data PromptData) (string, error) {
	// 解析模板
	tmpl, err := template.New("prompt").Parse(promptTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return "", err
	}

	// 执行模板
	var prompt bytes.Buffer
	err = tmpl.Execute(&prompt, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return "", err
	}
	return prompt.String(), nil
}
