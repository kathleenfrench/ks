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
	dKey = "decode"
	eKey = "encode"
)

func handleFile(t string) {
	if !strings.Contains(targetFile, "yaml") || strings.Contains(targetFile, "yml") {
		ui.ExitOnErr("target file must by YAML")
	}

	exists, err := fm.FilepathExists(targetFile)
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	if !exists {
		ui.ExitOnErr(fmt.Sprintf("%s does not exist - are you sure you provided the correct file path?", targetFile))
	}

	raw, _ := fm.ReadFile(targetFile)
	k8res, err := p.ParseK8sYAML(string(raw))
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	secretData := k8res.Blob.Object["data"]

	if verbose {
		raw, _ := fm.ReadFile(targetFile)
		theme.Info(fmt.Sprintf("--- %s ----", targetFile))
		fmt.Println(string(raw))
	}

	secretDataMap, err := p.InterfaceToMap(secretData)
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	var keys []string
	if secretData != nil {
		keys = p.GetMapKeys(secretDataMap)
	}

	if len(keys) == 0 {
		ui.ExitOnErr("no data keys to parse...")
	}

	var selected string
	prompt := &survey.Select{
		Message: "select an existing key",
		Options: keys,
	}

	err = survey.AskOne(prompt, &selected)
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	selectedValue := secretDataMap[selected]
	theme.Info(selectedValue)

	var selectedRoute string
	prompt = &survey.Select{
		Message: "do you want to decode or encode this value?",
		Options: []string{dKey, eKey, ui.QuitKey},
	}

	err = survey.AskOne(prompt, &selectedRoute)
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	if selectedRoute == ui.QuitKey {
		ui.Exit()
	}

	var nextStep string
	prompt = &survey.Select{
		Message: "select one",
		Options: []string{ui.CopyOnlyPromptMessage, ui.CopyAndOpenPromptMessage, ui.QuitKey},
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
		case ui.QuitKey:
			ui.Exit()
		default:
			err := decoder.Run(selectedValue, silent, verbose)
			if err != nil {
				ui.ExitOnErr(err.Error())
			}

			if nextStep == ui.CopyAndOpenPromptMessage {
				initCmd, err := ui.GetEditorPrompt("select an editor")
				if err != nil {
					ui.ExitOnErr(err.Error())
				}

				out, err := ui.GetTextEditorInputOnSave(fmt.Sprintf("view/edit %s", targetFile), string(fileContents), "**.yaml", initCmd)
				if err != nil {
					ui.ExitOnErr(err.Error())
				}

				err = fm.Write(targetFile, []byte(out))
				if err != nil {
					ui.ExitOnErr(err.Error())
				}

				if verbose {
					theme.Result(out)
				}

				theme.Info(fmt.Sprintf("saved changes to %s!", targetFile))
			}
		}
	case eKey:
		switch nextStep {
		case ui.QuitKey:
			ui.Exit()
		default:
			err := encoder.Run(selectedValue, silent, verbose)
			if err != nil {
				ui.ExitOnErr(err.Error())
			}

			if nextStep == ui.CopyAndOpenPromptMessage {
				initCmd, err := ui.GetEditorPrompt("select an editor")
				if err != nil {
					ui.ExitOnErr(err.Error())
				}

				out, err := ui.GetTextEditorInputOnSave(fmt.Sprintf("view/edit %s", targetFile), string(fileContents), "**.yaml", initCmd)
				if err != nil {
					ui.ExitOnErr(err.Error())
				}

				err = fm.Write(targetFile, []byte(out))
				if err != nil {
					ui.ExitOnErr(err.Error())
				}

				if verbose {
					theme.Result(out)
				}

				theme.Info(fmt.Sprintf("saved changes to %s!", targetFile))
			}
		}
	}
}
