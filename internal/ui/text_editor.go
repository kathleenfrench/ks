package ui

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kathleenfrench/ks/internal/theme"
)

const (
	VIM     = "Vim"
	EMACS   = "Emacs"
	ATOM    = "Atom"
	SUBLIME = "Sublime"
	VSCODE  = "VS Code"
)

var supportedEditors = []string{
	VIM,
	EMACS,
	ATOM,
	SUBLIME,
	VSCODE,
}

var editorInitMap = map[string]string{
	VIM:     "vim",
	EMACS:   "emacs",
	ATOM:    "atom --wait",
	VSCODE:  "code --wait",
	SUBLIME: "subl -n -w",
}

func GetEditorPrompt(msg string) (string, error) {
	supportedEditors = append(supportedEditors, "quit")

	var selected string
	prompt := &survey.Select{
		Message: msg,
		Options: supportedEditors,
	}

	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return selected, err
	}

	if selected == "quit" {
		theme.Info("bye!")
		os.Exit(0)
	}

	return editorInitMap[selected], nil
}

func GetTextEditorInputOnSave(msg string, defaultText string, filename string, editorInit string) (string, error) {
	var content string

	prompt := &survey.Editor{
		Message:       msg,
		FileName:      filename,
		Default:       defaultText,
		HideDefault:   true,
		AppendDefault: true,
		Editor:        editorInit,
		Help:          "save the file to close it",
	}

	err := survey.AskOne(prompt, &content)
	if err != nil {
		return content, err
	}

	return content, nil
}
