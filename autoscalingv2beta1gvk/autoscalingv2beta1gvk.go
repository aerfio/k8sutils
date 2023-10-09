package autoscalingv2beta1gvk

import (
	"reflect"

	"k8s.io/api/autoscaling/v2beta1"
)

var (
	HorizontalPodAutoscaler     = v2beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v2beta1.HorizontalPodAutoscaler{}).Name())
	HorizontalPodAutoscalerList = v2beta1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v2beta1.HorizontalPodAutoscalerList{}).Name())
)
