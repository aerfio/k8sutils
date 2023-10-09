package schedulingv1gvk

import (
	"reflect"

	"k8s.io/api/scheduling/v1"
)

var (
	PriorityClass     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PriorityClass{}).Name())
	PriorityClassList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PriorityClassList{}).Name())
)
