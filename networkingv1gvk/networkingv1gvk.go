package networkingv1gvk

import (
	"reflect"

	"k8s.io/api/networking/v1"
)

var (
	Ingress           = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Ingress{}).Name())
	IngressList       = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.IngressList{}).Name())
	IngressClass      = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.IngressClass{}).Name())
	IngressClassList  = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.IngressClassList{}).Name())
	NetworkPolicy     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.NetworkPolicy{}).Name())
	NetworkPolicyList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.NetworkPolicyList{}).Name())
)
