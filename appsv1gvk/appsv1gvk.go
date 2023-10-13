package appsv1gvk

import (
	"reflect"

	"k8s.io/api/apps/v1"
)

var (
	ControllerRevision     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ControllerRevision{}).Name())
	ControllerRevisionList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ControllerRevisionList{}).Name())
	DaemonSet              = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.DaemonSet{}).Name())
	DaemonSetList          = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.DaemonSetList{}).Name())
	Deployment             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Deployment{}).Name())
	DeploymentList         = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.DeploymentList{}).Name())
	ReplicaSet             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ReplicaSet{}).Name())
	ReplicaSetList         = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ReplicaSetList{}).Name())
	StatefulSet            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.StatefulSet{}).Name())
	StatefulSetList        = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.StatefulSetList{}).Name())
)
