package parse

import (
	"reflect"
)

func (p *parser) GetMapKeys(v map[string]string) (keys []string) {
	for k := range v {
		keys = append(keys, k)
	}

	return keys
}

func (p *parser) InterfaceToMap(v interface{}) (map[string]string, error) {
	res := make(map[string]string)
	iter := reflect.ValueOf(v).MapRange()
	for iter.Next() {
		key := iter.Key().Interface()
		value := iter.Value().Interface()
		res[key.(string)] = value.(string)
	}

	return res, nil
}
