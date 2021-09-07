package parse

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func (p *parser) CanParseYAML(content []byte, v interface{}) error {
	err := yaml.Unmarshal(content, &v)
	if err != nil {
		return fmt.Errorf("cannot parse yaml: %w", err)
	}

	return nil
}

func (p *parser) SupportedExtension(filepath string, exts []string) bool {
	for _, x := range exts {
		if strings.Contains(filepath, x) {
			return true
		}
	}

	return false
}
