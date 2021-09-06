package ks

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kathleenfrench/ks/internal/decoder"
	"github.com/kathleenfrench/ks/internal/encoder"
	"github.com/kathleenfrench/ks/internal/theme"
)

const (
	dCopyKey = "decode & copy to clipboard"
	eCopyKey = "encode & copy to clipboard"
	quitKey  = "exit"
)

func handleFile(t string) {
	if !strings.Contains(targetFile, "yaml") || strings.Contains(targetFile, "yml") {
		theme.Err("target file must be YAML")
		os.Exit(1)
	}

	exists, err := fm.FilepathExists(targetFile)
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	if !exists {
		theme.Err(fmt.Sprintf("%s does not exist - are you sure you provided the correct file path?", targetFile))
		os.Exit(1)
	}

	k8s, err := p.ReadSecretYAML(targetFile)
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	if verbose {
		raw, _ := fm.ReadFile(targetFile)
		theme.Info(fmt.Sprintf("--- %s ----", targetFile))
		fmt.Println(string(raw))
	}

	var keys []string
	if k8s.Data != nil {
		keys = p.GetMapKeys(k8s.Data)
	}

	if len(keys) == 0 {
		theme.Err("no data keys to parse...")
		os.Exit(1)
	}

	var selected string
	prompt := &survey.Select{
		Message: "select a key",
		Options: keys,
	}

	err = survey.AskOne(prompt, &selected)
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	selectedValue := k8s.Data[selected]
	theme.Info(selectedValue)

	var selectedAction string
	prompt = &survey.Select{
		Message: "select one",
		Options: []string{dCopyKey, eCopyKey, quitKey},
	}

	err = survey.AskOne(prompt, &selectedAction)
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	switch selectedAction {
	case dCopyKey:
		err := decoder.Run(selectedValue, silent, verbose)
		if err != nil {
			theme.Err(err.Error())
			os.Exit(1)
		}
	case eCopyKey:
		err := encoder.Run(selectedValue, silent, verbose)
		if err != nil {
			theme.Err(err.Error())
			os.Exit(1)
		}
	case quitKey:
		theme.Info("bye!")
		os.Exit(1)
	}
}
