package autoscalingv2beta2gvk

import (
	"reflect"

	"k8s.io/api/autoscaling/v2beta2"
)

var (
	HorizontalPodAutoscaler     = v2beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v2beta2.HorizontalPodAutoscaler{}).Name())
	HorizontalPodAutoscalerList = v2beta2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v2beta2.HorizontalPodAutoscalerList{}).Name())
)
