package coordinationv1beta1gvk

import (
	"reflect"

	"k8s.io/api/coordination/v1beta1"
)

var (
	Lease     = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.Lease{}).Name())
	LeaseList = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.LeaseList{}).Name())
)
