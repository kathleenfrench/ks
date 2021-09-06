package parse

import "encoding/base64"

func (p *parser) Encode(s string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	return string(encoded), nil
}
