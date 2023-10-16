package k8sutils

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func ToUnstructured(arg any) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	var err error
	obj.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(arg)
	return obj, err
}
