package secret

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kathleenfrench/ks/internal/ui"
	"github.com/kathleenfrench/ks/pkg/parse"

	yamlWriter "sigs.k8s.io/yaml"
)

func (m *manager) AddNewKey(b *parse.UnstructuredK8s, targetFile string) (*parse.UnstructuredK8s, error) {
	var (
		newKey         string
		newValue       string
		selectedAction string
		updatedData    = b.Data
		updatedKeys    = b.DataKeys
	)

	prompt := &survey.Input{
		Message: ui.ProvideValidSecretKeyMessage,
	}

	err := survey.AskOne(prompt, &newKey)
	if err != nil {
		return nil, err
	}

	if newKey == "" {
		return nil, errors.New("must provide a key")
	}

	newKey = strings.TrimSpace(newKey)
	updatedData[newKey] = ""
	updatedKeys = append(updatedKeys, newKey)

	prompt = &survey.Input{
		Message: ui.ProvideSecretValue,
	}

	err = survey.AskOne(prompt, &newValue)
	if err != nil {
		return nil, err
	}

	if newValue == "" {
		return nil, fmt.Errorf("must provide a value for %s", newKey)
	}

	newValue = strings.TrimSpace(newValue)

	sPrompt := &survey.Select{
		Message: ui.UpdateSecretBeforeSavingMessage,
		Options: []string{ui.EncodeKey, ui.DecodeKey, ui.SaveAsIs, ui.QuitKey},
	}

	err = survey.AskOne(sPrompt, &selectedAction)
	if err != nil {
		return nil, err
	}

	switch selectedAction {
	case ui.DecodeKey:
		decoded, err := m.parser.Decode(newValue)
		if err != nil {
			return nil, err
		}

		updatedData[newKey] = decoded
	case ui.EncodeKey:
		encoded, err := m.parser.Encode(newValue)
		if err != nil {
			return nil, err
		}

		updatedData[newKey] = encoded
	case ui.SaveAsIs:
		updatedData[newKey] = newValue
	case ui.QuitKey:
		ui.Exit()
	}

	b.Data = updatedData
	b.DataKeys = updatedKeys

	b.Blob.Object["data"] = updatedData

	yamlJSON, err := b.Blob.MarshalJSON()
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	jsonToYaml, err := yamlWriter.JSONToYAML(yamlJSON)
	if err != nil {
		ui.ExitOnErr(err.Error())
	}

	b.Raw = string(jsonToYaml)
	err = m.fm.Write(targetFile, []byte(b.Raw))
	if err != nil {
		return b, nil
	}

	return b, nil
}
