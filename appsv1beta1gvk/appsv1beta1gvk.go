package appsv1beta1gvk

import (
	"reflect"

	"k8s.io/api/apps/v1beta1"
)

var (
	StatefulSet            = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.StatefulSet{}).Name())
	StatefulSetList        = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.StatefulSetList{}).Name())
	ControllerRevision     = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.ControllerRevision{}).Name())
	ControllerRevisionList = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.ControllerRevisionList{}).Name())
	Deployment             = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.Deployment{}).Name())
	DeploymentList         = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.DeploymentList{}).Name())
	Scale                  = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.Scale{}).Name())
)
