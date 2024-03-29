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
	objectIdentityExtactor
}

var _ client.SubResourceClient = &K8sSubresourceClient{}

func (k *K8sSubresourceClient) Get(ctx context.Context, obj, subResource client.Object, opts ...client.SubResourceGetOption) error {
	sctx, span := k.tracer.Start(ctx, "Get")
	defer span.End()

	getOpts := &client.SubResourceGetOptions{}
	for _, opt := range opts {
		opt.ApplyToSubResourceGet(getOpts)
	}

	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)
	k.objectIdentityExtactor.handleClientObjectAttrs(span, true, subResource)

	handleGetOptions(span, &client.GetOptions{Raw: getOpts.AsGetOptions()})

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

	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)
	k.objectIdentityExtactor.handleClientObjectAttrs(span, true, subResource)
	subResOpt := &client.SubResourceCreateOptions{}
	for _, opt := range opts {
		opt.ApplyToSubResourceCreate(subResOpt)
	}
	handleCreateOptions(span, &client.CreateOptions{
		Raw: subResOpt.AsCreateOptions(),
	})

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
	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)
	subResOpt := &client.SubResourceUpdateOptions{}
	for _, opt := range opts {
		opt.ApplyToSubResourceUpdate(subResOpt)
	}
	handleUpdateOpts(span, subResOpt.AsUpdateOptions())

	if err := k.inner.Update(sctx, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to create a subresource")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
	}

	return nil
}

func (k *K8sSubresourceClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	sctx, span := k.tracer.Start(ctx, "Patch")
	defer span.End()

	span.SetAttributes(attribute.String("patch.type", string(patch.Type())))
	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)

	subResOpt := &client.SubResourcePatchOptions{}
	for _, opt := range opts {
		opt.ApplyToSubResourcePatch(subResOpt)
	}
	// this ignores subResourceBody field of SubResourcePatchOptions, but it's on purpose, it would be too big to fit into span attribute
	handlePatchOpts(span, subResOpt.AsPatchOptions())

	if err := k.inner.Patch(sctx, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to create a subresource")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
	}

	return nil
}
