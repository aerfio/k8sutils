package batchv1gvk

import (
	"reflect"

	"k8s.io/api/batch/v1"
)

var (
	CronJob     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CronJob{}).Name())
	CronJobList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.CronJobList{}).Name())
	Job         = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Job{}).Name())
	JobList     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.JobList{}).Name())
)
