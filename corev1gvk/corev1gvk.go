package corev1gvk

import (
	"reflect"

	"k8s.io/api/core/v1"
)

var (
	Binding                   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Binding{}).Name())
	ConfigMap                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ConfigMap{}).Name())
	ConfigMapList             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ConfigMapList{}).Name())
	Endpoints                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Endpoints{}).Name())
	EndpointsList             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.EndpointsList{}).Name())
	Event                     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Event{}).Name())
	EventList                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.EventList{}).Name())
	LimitRange                = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.LimitRange{}).Name())
	LimitRangeList            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.LimitRangeList{}).Name())
	Namespace                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Namespace{}).Name())
	NamespaceList             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.NamespaceList{}).Name())
	Node                      = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Node{}).Name())
	NodeList                  = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.NodeList{}).Name())
	PersistentVolume          = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolume{}).Name())
	PersistentVolumeList      = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolumeList{}).Name())
	PersistentVolumeClaim     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolumeClaim{}).Name())
	PersistentVolumeClaimList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolumeClaimList{}).Name())
	Pod                       = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Pod{}).Name())
	PodList                   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodList{}).Name())
	PodTemplate               = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodTemplate{}).Name())
	PodTemplateList           = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodTemplateList{}).Name())
	RangeAllocation           = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.RangeAllocation{}).Name())
	ReplicationController     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ReplicationController{}).Name())
	ReplicationControllerList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ReplicationControllerList{}).Name())
	ResourceQuota             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ResourceQuota{}).Name())
	ResourceQuotaList         = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ResourceQuotaList{}).Name())
	Secret                    = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Secret{}).Name())
	SecretList                = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.SecretList{}).Name())
	Service                   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Service{}).Name())
	ServiceList               = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ServiceList{}).Name())
	ServiceAccount            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ServiceAccount{}).Name())
	ServiceAccountList        = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ServiceAccountList{}).Name())
)
