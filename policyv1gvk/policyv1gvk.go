package policyv1gvk

import (
	"reflect"

	"k8s.io/api/policy/v1"
)

var (
	PodDisruptionBudget     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodDisruptionBudget{}).Name())
	PodDisruptionBudgetList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.PodDisruptionBudgetList{}).Name())
)
