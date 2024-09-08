package prompt

import (
	"bytes"
	"fmt"
	"golang.org/x/text/language"
	"os"
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

func NewPromptData(database, content string) PromptData {
	return PromptData{
		DatabaseName: database,
		SlowQueryLog: content,
		Language:     getSystemLanguage(),
	}
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

func getSystemLanguage() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		lang = os.Getenv("LANGUAGE")
	}
	tag, err := language.Parse(lang)
	if err != nil {
		return "zh-CN"
	}
	return tag.String()
}
