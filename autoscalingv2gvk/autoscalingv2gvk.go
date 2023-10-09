package autoscalingv2gvk

import (
	"reflect"

	"k8s.io/api/autoscaling/v2"
)

var (
	HorizontalPodAutoscaler     = v2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v2.HorizontalPodAutoscaler{}).Name())
	HorizontalPodAutoscalerList = v2.SchemeGroupVersion.WithKind(reflect.TypeOf(&v2.HorizontalPodAutoscalerList{}).Name())
)
