package parse

import "encoding/base64"

func (p *parser) Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return s, err
	}

	return string(decoded), nil
}
