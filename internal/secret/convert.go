package secret

import (
	"fmt"

	"github.com/kathleenfrench/ks/pkg/parse"
	yamlWriter "sigs.k8s.io/yaml"
)

func (m *manager) EncodeData(b *parse.UnstructuredK8s) (*parse.UnstructuredK8s, error) {
	encodedValues := make(map[string]string)

	for k, v := range b.Data {
		encoded, err := m.parser.Encode(v)
		if err != nil {
			return nil, fmt.Errorf("could not encode %s: %w", k, err)
		}

		encodedValues[k] = encoded
	}

	b.Data = encodedValues
	b.Blob.Object["data"] = b.Data

	yamlJSON, err := b.Blob.MarshalJSON()
	if err != nil {
		return nil, err
	}

	jsonToYaml, err := yamlWriter.JSONToYAML(yamlJSON)
	if err != nil {
		return nil, err
	}

	b.Raw = string(jsonToYaml)

	return b, nil
}

func (m *manager) DecodeData(b *parse.UnstructuredK8s) (*parse.UnstructuredK8s, error) {
	decodedValues := make(map[string]string)

	for k, v := range b.Data {
		decoded, err := m.parser.Decode(v)
		if err != nil {
			return nil, fmt.Errorf("could not decode %s: %w", k, err)
		}

		decodedValues[k] = decoded
	}

	b.Data = decodedValues
	b.Blob.Object["data"] = b.Data

	yamlJSON, err := b.Blob.MarshalJSON()
	if err != nil {
		return nil, err
	}

	jsonToYaml, err := yamlWriter.JSONToYAML(yamlJSON)
	if err != nil {
		return nil, err
	}

	b.Raw = string(jsonToYaml)

	return b, nil
}
