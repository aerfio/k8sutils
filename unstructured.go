package k8sutils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func ToUnstructured(arg any) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	var err error
	obj.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(arg)
	return obj, err
}

type ObjectList interface {
	metav1.ListInterface
	runtime.Object
}
type Object[T any] interface {
	*T
	metav1.Object
	runtime.Object
}

func GetListObjects[V Object[T], T any](list ObjectList) ([]V, error) {
	unstrList := &unstructured.UnstructuredList{}
	raw, err := runtime.DefaultUnstructuredConverter.ToUnstructured(list)
	if err != nil {
		return nil, err
	}
	unstrList.SetUnstructuredContent(raw)
	objList := make([]V, 0, len(unstrList.Items))
	for _, obj := range unstrList.Items {
		out := V(new(T))
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), out); err != nil {
			return nil, err
		}
		objList = append(objList, out)
	}
	return objList, nil
}
