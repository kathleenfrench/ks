package ks

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/kathleenfrench/ks/internal/secret"
	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/internal/ui"
)

func handleFile(t string) {
	var (
		keys     []string
		selected string
		sm       = secret.NewManager()
	)

	blob, err := sm.Parse(targetFile)
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	if verbose {
		sm.PrintFile(targetFile, blob.Raw)
	}

	if blob.Data != nil {
		keys = append(blob.DataKeys, ui.AddNewKey)
	}

	prompt := &survey.Select{
		Message: ui.SelectFromAnExistingKeyMessage,
		Options: keys,
	}

	err = survey.AskOne(prompt, &selected)
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	switch selected {
	case ui.AddNewKey:
		updatedYaml, err := sm.AddNewKey(blob)
		if err != nil {
			ui.ExitOnErr(err.Error())
		}

		if verbose && !silent {
			theme.Result(updatedYaml)
		}
	default:
		err = sm.Handle(blob, blob.Data[selected], targetFile, silent, verbose)
		if err != nil {
			ui.ExitOnErr(err.Error())
		}
	}
}
