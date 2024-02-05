package k8sutils

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

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

func AddTypeInformationToObject(scheme *runtime.Scheme, obj runtime.Object) error {
	gvk, err := apiutil.GVKForObject(obj, scheme)
	if err != nil {
		return fmt.Errorf("failed to get GVK for %T: %w", obj, err)
	}
	obj.GetObjectKind().SetGroupVersionKind(gvk)
	return nil
}
