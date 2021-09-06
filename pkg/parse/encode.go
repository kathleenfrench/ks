package parse

import "encoding/base64"

func (p *parser) Encode(s string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	return string(encoded), nil
}

func (p *parser) Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return s, err
	}

	return string(decoded), nil
}
