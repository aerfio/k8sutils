package tracing

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SetSpanErr(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}

type objectIdentityExtactor struct {
	groupVersionKindForFn func(object runtime.Object) (schema.GroupVersionKind, error)
}

func (o *objectIdentityExtactor) handleClientObjectAttrs(span trace.Span, isSubResource bool, obj client.Object) {
	objectType := "object"
	if isSubResource {
		objectType = "subresource"
	}

	if ns := obj.GetNamespace(); ns != "" {
		span.SetAttributes(attribute.String("object.namespace", ns))
	}
	span.SetAttributes(attribute.String("object.name", obj.GetName()))

	gvk, err := o.groupVersionKindForFn(obj)
	if err != nil {
		span.RecordError(fmt.Errorf("failed to get GVK for object: %s", err))
	} else {
		span.SetAttributes(
			attribute.String(fmt.Sprintf("%s.group", objectType), gvk.Group),
			attribute.String(fmt.Sprintf("%s.version", objectType), gvk.Version),
			attribute.String(fmt.Sprintf("%s.kind", objectType), gvk.Kind),
		)
	}
}
