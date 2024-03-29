package tracing

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type K8sSubresourceClient struct {
	inner                 client.SubResourceClient
	groupVersionKindForFn func(object runtime.Object) (schema.GroupVersionKind, error)
	tracer                trace.Tracer
}

var _ client.SubResourceClient = &K8sSubresourceClient{}

func (k *K8sSubresourceClient) handleObjectAttrs(span trace.Span, obj client.Object) {
	gvk, err := k.groupVersionKindForFn(obj)
	if err != nil {
		span.AddEvent("failed to get GVK for object", trace.WithAttributes(attribute.String("error", err.Error())))
	} else {
		span.SetAttributes(
			attribute.String("object.group", gvk.Group),
			attribute.String("object.version", gvk.Version),
			attribute.String("object.kind", gvk.Kind),
		)
	}
}

func (k *K8sSubresourceClient) handleObjectAndSubresourceAttrs(span trace.Span, obj, subResource client.Object) {
	gvk, err := k.groupVersionKindForFn(obj)
	if err != nil {
		span.AddEvent("failed to get GVK for object", trace.WithAttributes(attribute.String("error", err.Error())))
	} else {
		span.SetAttributes(
			attribute.String("object.group", gvk.Group),
			attribute.String("object.version", gvk.Version),
			attribute.String("object.kind", gvk.Kind),
		)
	}

	subResourceGVK, err := k.groupVersionKindForFn(subResource)
	if err != nil {
		span.AddEvent("failed to get GVK for subresource", trace.WithAttributes(attribute.String("error", err.Error())))
	} else {
		span.SetAttributes(
			attribute.String("subresource.group", subResourceGVK.Group),
			attribute.String("subresource.version", subResourceGVK.Version),
			attribute.String("subresource.kind", subResourceGVK.Kind),
		)
	}
}

func (k *K8sSubresourceClient) Get(ctx context.Context, obj, subResource client.Object, opts ...client.SubResourceGetOption) error {
	sctx, span := k.tracer.Start(ctx, "Get")
	defer span.End()

	if name := obj.GetName(); name != "" {
		span.SetAttributes(attribute.String("object.name", name))
	}
	if ns := obj.GetNamespace(); ns != "" {
		span.SetAttributes(attribute.String("object.namespace", ns))
	}
	getOpts := &client.SubResourceGetOptions{}
	for _, opt := range opts {
		opt.ApplyToSubResourceGet(getOpts)
	}

	handleGetOptions(span, &client.GetOptions{Raw: getOpts.AsGetOptions()})
	k.handleObjectAndSubresourceAttrs(span, obj, subResource)

	if err := k.inner.Get(sctx, obj, subResource, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to get a subresource")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
	}

	return nil
}

func (k *K8sSubresourceClient) Create(ctx context.Context, obj, subResource client.Object, opts ...client.SubResourceCreateOption) error {
	sctx, span := k.tracer.Start(ctx, "Create")
	defer span.End()
	if name := obj.GetName(); name != "" {
		span.SetAttributes(attribute.String("object.name", name))
	}
	if ns := obj.GetNamespace(); ns != "" {
		span.SetAttributes(attribute.String("object.namespace", ns))
	}
	subResOpt := &client.SubResourceCreateOptions{}
	for _, opt := range opts {
		opt.ApplyToSubResourceCreate(subResOpt)
	}
	handleCreateOptions(span, &client.CreateOptions{
		Raw: subResOpt.AsCreateOptions(),
	})

	k.handleObjectAndSubresourceAttrs(span, obj, subResource)

	if err := k.inner.Create(sctx, obj, subResource, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to create a subresource")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
	}

	return nil
}

func (k *K8sSubresourceClient) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	sctx, span := k.tracer.Start(ctx, "Update")
	defer span.End()
	if name := obj.GetName(); name != "" {
		span.SetAttributes(attribute.String("object.name", name))
	}
	if ns := obj.GetNamespace(); ns != "" {
		span.SetAttributes(attribute.String("object.namespace", ns))
	}
	subResOpt := &client.SubResourceUpdateOptions{}
	for _, opt := range opts {
		opt.ApplyToSubResourceUpdate(subResOpt)
	}
	handleUpdateOpts(span, subResOpt.AsUpdateOptions())

	k.handleObjectAttrs(span, obj)

	if err := k.inner.Update(sctx, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to create a subresource")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
	}

	return nil
}

func (k *K8sSubresourceClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	// TODO implement me
	panic("implement me")
}
