package coordinationv1gvk

import (
	"reflect"

	"k8s.io/api/coordination/v1"
)

var (
	Lease     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Lease{}).Name())
	LeaseList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.LeaseList{}).Name())
)
