package corev1gvk

import (
	"reflect"

	"k8s.io/api/core/v1"
)

var (
	ReplicationController     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ReplicationController{}).Name())
	ReplicationControllerList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ReplicationControllerList{}).Name())
	Binding                   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Binding{}).Name())
	Node                      = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Node{}).Name())
	NodeList                  = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.NodeList{}).Name())
	Service                   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Service{}).Name())
	ServiceList               = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ServiceList{}).Name())
	PersistentVolume          = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolume{}).Name())
	PersistentVolumeList      = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolumeList{}).Name())
	Secret                    = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Secret{}).Name())
	SecretList                = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.SecretList{}).Name())
	RangeAllocation           = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.RangeAllocation{}).Name())
	ResourceQuota             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ResourceQuota{}).Name())
	ResourceQuotaList         = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ResourceQuotaList{}).Name())
	PersistentVolumeClaim     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolumeClaim{}).Name())
	PersistentVolumeClaimList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PersistentVolumeClaimList{}).Name())
	Event                     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Event{}).Name())
	EventList                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.EventList{}).Name())
	PodTemplate               = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodTemplate{}).Name())
	PodTemplateList           = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodTemplateList{}).Name())
	ServiceAccount            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ServiceAccount{}).Name())
	ServiceAccountList        = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ServiceAccountList{}).Name())
	ConfigMap                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ConfigMap{}).Name())
	ConfigMapList             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ConfigMapList{}).Name())
	Endpoints                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Endpoints{}).Name())
	EndpointsList             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.EndpointsList{}).Name())
	LimitRange                = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.LimitRange{}).Name())
	LimitRangeList            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.LimitRangeList{}).Name())
	Namespace                 = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Namespace{}).Name())
	NamespaceList             = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.NamespaceList{}).Name())
	Pod                       = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Pod{}).Name())
	PodList                   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodList{}).Name())
)
