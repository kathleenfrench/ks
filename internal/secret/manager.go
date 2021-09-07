package secret

import (
	"fmt"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/pkg/file"
	"github.com/kathleenfrench/ks/pkg/parse"
)

type Manager interface {
	Handle(b *parse.UnstructuredK8s, selectedValue string, targetFilename string, silent bool, verbose bool) error
	Parse(filepath string) (*parse.UnstructuredK8s, error)
	AddNewKey(b *parse.UnstructuredK8s, targetFile string) (*parse.UnstructuredK8s, error)
	PrintFile(name string, content string)
	EncodeData(b *parse.UnstructuredK8s) (*parse.UnstructuredK8s, error)
	DecodeData(b *parse.UnstructuredK8s) (*parse.UnstructuredK8s, error)
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

func (m *manager) PrintFile(name string, content string) {
	theme.Info(fmt.Sprintf("--- %s ----", name))
	fmt.Println(content)
}
