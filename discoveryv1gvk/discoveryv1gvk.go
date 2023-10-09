package discoveryv1gvk

import (
	"reflect"

	"k8s.io/api/discovery/v1"
)

var (
	EndpointSlice     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.EndpointSlice{}).Name())
	EndpointSliceList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.EndpointSliceList{}).Name())
)
