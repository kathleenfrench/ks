package secret

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kathleenfrench/ks/internal/decoder"
	"github.com/kathleenfrench/ks/internal/encoder"
	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/internal/ui"
	"github.com/kathleenfrench/ks/pkg/parse"
)

func (m *manager) Parse(filepath string) (*parse.UnstructuredK8s, error) {
	if !m.parser.SupportedExtension(filepath, []string{".yaml", ".yml"}) {
		return nil, fmt.Errorf("%s does not have a valid yaml or yml extension", filepath)
	}

	fileExists, err := m.fm.FilepathExists(filepath)
	if err != nil {
		return nil, err
	}

	if !fileExists {
		return nil, fmt.Errorf("%s does not exist - are you sure you provided the correct file path?", filepath)
	}

	raw, err := m.fm.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	unstructuredOutput, err := m.parser.ParseK8sYAML(string(raw))
	if err != nil {
		return nil, err
	}

	return unstructuredOutput, nil
}

func (m *manager) Handle(b *parse.UnstructuredK8s, selectedValue string, targetFilename string, silent bool, verbose bool) error {
	if !silent {
		theme.Info(selectedValue)
	}

	prompt := &survey.Select{
		Message: ui.DecodeOrEncodePromptMessage,
		Options: []string{ui.DecodeKey, ui.EncodeKey, ui.QuitKey},
	}

	var selectedConversion string
	err := survey.AskOne(prompt, &selectedConversion)
	if err != nil {
		return err
	}

	if selectedConversion == ui.QuitKey {
		ui.Exit()
	}

	prompt = &survey.Select{
		Message: ui.SelectNextActionMessage,
		Options: []string{ui.CopyOnlyPromptMessage, ui.CopyAndOpenPromptMessage, ui.QuitKey},
	}

	var selectedOp string
	err = survey.AskOne(prompt, &selectedOp)
	if err != nil {
		return err
	}

	if selectedOp == ui.QuitKey {
		ui.Exit()
	}

	switch selectedConversion {
	case ui.DecodeKey:
		err := decoder.Run(selectedValue, silent)
		if err != nil {
			ui.ExitOnErr(err.Error())
		}
	case ui.EncodeKey:
		err := encoder.Run(selectedValue, silent)
		if err != nil {
			ui.ExitOnErr(err.Error())
		}
	}

	if selectedOp != ui.CopyAndOpenPromptMessage {
		return nil
	}

	initEditorCmd, err := ui.GetEditorPrompt(ui.SelectAnEditor)
	if err != nil {
		return err
	}

	output, err := ui.GetTextEditorInputOnSave(fmt.Sprintf("View and/or Edit %s", targetFilename), b.Raw, "**.yaml", initEditorCmd)
	if err != nil {
		return err
	}

	err = m.fm.Write(targetFilename, []byte(output))
	if err != nil {
		return err
	}

	if !silent {
		if verbose {
			m.PrintFile(targetFilename, output)
		}

		theme.Info(fmt.Sprintf("saved any changes to %s!", targetFilename))
	}

	return nil
}
