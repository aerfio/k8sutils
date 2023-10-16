package k8sutils_test

import (
	"strings"
	"testing"

	corev1 "k8s.io/api/core/v1"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"

	"aerf.io/k8sutils"
)

func TestGetListObjects(t *testing.T) {
	configMapList := &corev1.ConfigMapList{}
	decoder := yamlutil.NewYAMLOrJSONDecoder(strings.NewReader(configMapListFixture), 2048)
	if err := decoder.Decode(configMapList); err != nil {
		t.Fatalf("failed to decode: %s", err)
	}

	configMaps, err := k8sutils.GetListObjects[*corev1.ConfigMap](configMapList)
	if err != nil {
		t.Fatalf("failed to convert &corev1.ConfigMapList{} to []*corev1.Configmap{}: %s", err)
	}
	if len(configMaps) != 2 {
		t.Errorf("there should be 2 elements in that list, we got %d, list: %#v", len(configMaps), configMaps)
	}
}

const configMapListFixture = `{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "v1",
            "data": {
                "key1": "config1",
                "key2": "config2"
            },
            "kind": "ConfigMap",
            "metadata": {
                "creationTimestamp": "2023-10-13T10:32:29Z",
                "name": "my-config",
                "namespace": "aerfio",
                "resourceVersion": "500192628",
                "uid": "88a7b4f5-4efd-4541-a9ff-7e62309d8abd"
            }
        },
        {
            "apiVersion": "v1",
            "data": {
                "key1": "config1",
                "key2": "config2"
            },
            "kind": "ConfigMap",
            "metadata": {
                "creationTimestamp": "2023-10-13T10:32:33Z",
                "name": "my-config2",
                "namespace": "aerfio",
                "resourceVersion": "500192681",
                "uid": "619b12b5-8ee9-463b-9495-ce2ff50d6dd2"
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": ""
    }
}
`
