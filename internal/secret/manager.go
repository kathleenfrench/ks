package secret

import (
	"github.com/kathleenfrench/ks/pkg/file"
	"github.com/kathleenfrench/ks/pkg/parse"
	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
)

type Manager interface {
	Parse(content []byte) (*v1.Secret, error)
}

type manager struct {
	parser parse.Parser
	fm     file.Manager
}

func NewManager() Manager {
	return &manager{
		parser: parse.NewParser(),
		fm:     file.NewManager(),
	}
}

func (m *manager) Parse(content []byte) (*v1.Secret, error) {
	s := &v1.Secret{}
	err := yaml.Unmarshal(content, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
