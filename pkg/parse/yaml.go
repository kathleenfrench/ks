package parse

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func (p *parser) CanParseYAML(content []byte, v interface{}) error {
	err := yaml.Unmarshal(content, &v)
	if err != nil {
		return fmt.Errorf("cannot parse yaml: %w", err)
	}

	return nil
}
