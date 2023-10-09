package batchv1beta1gvk

import (
	"reflect"

	"k8s.io/api/batch/v1beta1"
)

var (
	CronJob     = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.CronJob{}).Name())
	CronJobList = v1beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1beta1.CronJobList{}).Name())
)
