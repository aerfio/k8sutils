package autoscalingv1gvk

import (
	"reflect"

	"k8s.io/api/autoscaling/v1"
)

var (
	HorizontalPodAutoscaler     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.HorizontalPodAutoscaler{}).Name())
	HorizontalPodAutoscalerList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.HorizontalPodAutoscalerList{}).Name())
	Scale                       = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Scale{}).Name())
)
