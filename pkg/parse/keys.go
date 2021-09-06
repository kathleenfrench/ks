package parse

func (p *parser) GetMapKeys(v map[string]string) (keys []string) {
	for k := range v {
		keys = append(keys, k)
	}

	return keys
}
