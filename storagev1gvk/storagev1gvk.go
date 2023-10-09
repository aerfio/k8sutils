package storagev1gvk

import (
	"reflect"

	"k8s.io/api/storage/v1"
)

var (
	StorageClass           = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.StorageClass{}).Name())
	StorageClassList       = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.StorageClassList{}).Name())
	CSIStorageCapacity     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CSIStorageCapacity{}).Name())
	CSIStorageCapacityList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CSIStorageCapacityList{}).Name())
	CSIDriver              = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CSIDriver{}).Name())
	CSIDriverList          = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CSIDriverList{}).Name())
	VolumeAttachment       = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.VolumeAttachment{}).Name())
	VolumeAttachmentList   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.VolumeAttachmentList{}).Name())
	CSINode                = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CSINode{}).Name())
	CSINodeList            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CSINodeList{}).Name())
)
