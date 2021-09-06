package parse

import "github.com/kathleenfrench/ks/pkg/file"

type Parser interface {
	Encode(s string) (string, error)
	Decode(s string) (string, error)
	ReadSecretYAML(filepath string) (*K8sMetaSecret, error)
	GetMapKeys(v map[string]string) (keys []string)
}

type parser struct {
	fm file.Manager
}

func NewParser() Parser {
	return &parser{
		fm: file.NewManager(),
	}
}
