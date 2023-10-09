package eventsv1gvk

import (
	"reflect"

	"k8s.io/api/events/v1"
)

var (
	Event     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Event{}).Name())
	EventList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.EventList{}).Name())
)
