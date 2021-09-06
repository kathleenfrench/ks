package parse

import (
	"encoding/json"

	"github.com/icza/dyno"
	"gopkg.in/yaml.v3"
	metav1_unstruct "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type UnstructuredK8s struct {
	Blob *metav1_unstruct.Unstructured
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
	}

	return out, nil
}
