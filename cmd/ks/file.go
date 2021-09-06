package ks

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kathleenfrench/ks/internal/decoder"
	"github.com/kathleenfrench/ks/internal/encoder"
	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/internal/ui"
)

const (
	dKey    = "decode"
	eKey    = "encode"
	quitKey = "exit"
)

const (
	copyOnlyKey    = "copy only"
	copyAndOpenKey = "copy & open target file"
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

	var selectedRoute string
	prompt = &survey.Select{
		Message: "do you want to decode or encode this value?",
		Options: []string{dKey, eKey, quitKey},
	}

	err = survey.AskOne(prompt, &selectedRoute)
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	if selectedRoute == quitKey {
		theme.Info("bye!")
		os.Exit(1)
	}

	var nextStep string
	prompt = &survey.Select{
		Message: "select one",
		Options: []string{copyOnlyKey, copyAndOpenKey, quitKey},
	}

	err = survey.AskOne(prompt, &nextStep)
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	fileContents, _ := fm.ReadFile(targetFile)

	switch selectedRoute {
	case dKey:
		switch nextStep {
		case quitKey:
			theme.Info("bye!")
			os.Exit(1)
		default:
			err := decoder.Run(selectedValue, silent, verbose)
			if err != nil {
				theme.Err(err.Error())
				os.Exit(1)
			}

			if nextStep == copyAndOpenKey {
				initCmd, err := ui.GetEditorPrompt("select an editor")
				if err != nil {
					theme.Err(err.Error())
					os.Exit(1)
				}

				out, err := ui.GetTextEditorInputOnSave(fmt.Sprintf("view/edit %s", targetFile), string(fileContents), "**.yaml", initCmd)
				if err != nil {
					theme.Err(err.Error())
					os.Exit(1)
				}

				err = fm.Write(targetFile, []byte(out))
				if err != nil {
					theme.Err(err.Error())
					os.Exit(1)
				}

				if verbose {
					theme.Result(out)
				}

				theme.Info(fmt.Sprintf("saved changes to %s!", targetFile))
			}
		}
	case eKey:
		switch nextStep {
		case quitKey:
			theme.Info("bye!")
			os.Exit(1)
		default:
			err := encoder.Run(selectedValue, silent, verbose)
			if err != nil {
				theme.Err(err.Error())
				os.Exit(1)
			}

			if nextStep == copyAndOpenKey {
				initCmd, err := ui.GetEditorPrompt("select an editor")
				if err != nil {
					theme.Err(err.Error())
					os.Exit(1)
				}

				out, err := ui.GetTextEditorInputOnSave(fmt.Sprintf("view/edit %s", targetFile), string(fileContents), "**.yaml", initCmd)
				if err != nil {
					theme.Err(err.Error())
					os.Exit(1)
				}

				err = fm.Write(targetFile, []byte(out))
				if err != nil {
					theme.Err(err.Error())
					os.Exit(1)
				}

				if verbose {
					theme.Result(out)
				}

				theme.Info(fmt.Sprintf("saved changes to %s!", targetFile))
			}
		}
	}
}
