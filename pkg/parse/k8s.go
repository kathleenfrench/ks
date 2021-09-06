package parse

import "gopkg.in/yaml.v3"

type K8sMetaSecret struct {
	Version  string            `yaml:"apiVersion"`
	Kind     string            `yaml:"kind"`
	Metadata map[string]string `yaml:"metadata"`
	Data     map[string]string `yaml:"data"`
	Type     string            `yaml:"type"`
}

func (p *parser) ReadSecretYAML(filepath string) (*K8sMetaSecret, error) {
	var secretYaml *K8sMetaSecret

	rawFile, err := p.fm.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(rawFile, &secretYaml)
	if err != nil {
		return nil, err
	}

	return secretYaml, nil
}
