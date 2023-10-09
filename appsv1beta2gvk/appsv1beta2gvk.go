package appsv1beta2gvk

import (
	"reflect"

	"k8s.io/api/apps/v1beta2"
)

var (
	ReplicaSet             = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.ReplicaSet{}).Name())
	ReplicaSetList         = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.ReplicaSetList{}).Name())
	ControllerRevision     = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.ControllerRevision{}).Name())
	ControllerRevisionList = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.ControllerRevisionList{}).Name())
	Deployment             = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.Deployment{}).Name())
	DeploymentList         = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.DeploymentList{}).Name())
	Scale                  = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.Scale{}).Name())
	StatefulSet            = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.StatefulSet{}).Name())
	StatefulSetList        = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.StatefulSetList{}).Name())
	DaemonSet              = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.DaemonSet{}).Name())
	DaemonSetList          = v1beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta2.DaemonSetList{}).Name())
)
