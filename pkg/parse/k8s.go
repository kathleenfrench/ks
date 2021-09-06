package parse

import (
	"encoding/json"

	"github.com/icza/dyno"
	"gopkg.in/yaml.v3"
	metav1_unstruct "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type UnstructuredK8s struct {
	Blob     *metav1_unstruct.Unstructured
	Data     map[string]string
	DataKeys []string
	Raw      string
}

func (p *parser) ParseK8sYAML(content string) (*UnstructuredK8s, error) {
	raw := &map[string]interface{}{}
	err := yaml.Unmarshal([]byte(content), raw)
	if err != nil {
		return nil, err
	}

	k8sJSON, err := json.Marshal(dyno.ConvertMapI2MapS(*raw))
	if err != nil {
		return nil, err
	}

	unstructured := metav1_unstruct.Unstructured{}
	err = unstructured.UnmarshalJSON(k8sJSON)
	if err != nil {
		return nil, err
	}

	out := &UnstructuredK8s{
		Blob: &unstructured,
		Raw:  content,
	}

	dataInterface, ok := unstructured.Object["data"]
	if !ok || dataInterface == nil {
		return out, nil
	}

	dataMap, err := p.InterfaceToMap(dataInterface)
	if err != nil {
		return out, err
	}

	out.Data = dataMap
	out.DataKeys = p.GetMapKeys(dataMap)

	return out, nil
}
