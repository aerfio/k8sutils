package k8stracing

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

func (o *objectIdentityExtactor) handleRuntimeObjectAttrs(span trace.Span, objectType string, obj runtime.Object) {
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

func (o *objectIdentityExtactor) handleClientObjectAttrs(span trace.Span, isSubResource bool, obj client.Object) {
	objectType := "object"
	if isSubResource {
		objectType = "subresource"
	}
	o.handleRuntimeObjectAttrs(span, objectType, obj)

	if ns := obj.GetNamespace(); ns != "" {
		span.SetAttributes(attribute.String(fmt.Sprintf("%s.namespace", objectType), ns))
	}
	if genName := obj.GetGenerateName(); genName != "" {
		span.SetAttributes(attribute.String(fmt.Sprintf("%s.name", objectType), obj.GetName()))
	}
	if name := obj.GetName(); name != "" {
		span.SetAttributes(attribute.String(fmt.Sprintf("%s.name", objectType), obj.GetName()))
	}
}

func (o *objectIdentityExtactor) handleClientObjectListAttrs(span trace.Span, obj client.ObjectList) {
	o.handleRuntimeObjectAttrs(span, "list", obj)

	if cont := obj.GetContinue(); cont != "" {
		span.SetAttributes(attribute.String("list.continue", cont))
	}
	if resourceVer := obj.GetResourceVersion(); resourceVer != "" {
		span.SetAttributes(attribute.String("list.resourceVersion", resourceVer))
	}
	if ric := obj.GetRemainingItemCount(); ric != nil {
		span.SetAttributes(attribute.Int64("list.remainingItemCount", *ric))
	}
}
