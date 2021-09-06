package parse

import (
	"github.com/kathleenfrench/ks/pkg/file"
)

type Parser interface {
	Encode(s string) (string, error)
	Decode(s string) (string, error)
	GetMapKeys(v map[string]string) (keys []string)
	CanParseYAML(content []byte, v interface{}) error
	ParseK8sYAML(content string) (*UnstructuredK8s, error)
	InterfaceToMap(v interface{}) (map[string]string, error)
	SupportedExtension(filepath string, exts []string) bool
}

type parser struct {
	fm file.Manager
}

func NewParser() Parser {
	return &parser{
		fm: file.NewManager(),
	}
}
