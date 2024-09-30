package k8stracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type K8sClient struct {
	inner                  client.Client
	tracer                 trace.Tracer
	traceProvider          trace.TracerProvider
	objectIdentityExtactor objectIdentityExtactor
}

func NewK8sClient(inner client.Client, traceProvider trace.TracerProvider) client.Client {
	return &K8sClient{
		inner:         inner,
		tracer:        traceProvider.Tracer("KubernetesClient"),
		traceProvider: traceProvider,
		objectIdentityExtactor: objectIdentityExtactor{
			groupVersionKindForFn: inner.GroupVersionKindFor,
		},
	}
}

func handleGetOptions(span trace.Span, opts ...client.GetOption) {
	if len(opts) == 0 {
		return
	}

	getOpts := &client.GetOptions{}

	getOpts.ApplyOptions(opts)
	metav1GetOpts := getOpts.AsGetOptions()
	if resVer := metav1GetOpts.ResourceVersion; resVer != "" {
		span.SetAttributes(attribute.String("getOption.resourceVersion", resVer))
	}
}

func (k *K8sClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	sctx, span := k.tracer.Start(ctx, "Get", trace.WithAttributes(attribute.String("name", key.Name), attribute.String("namespace", key.Namespace)))
	defer span.End()

	getOpts := &client.GetOptions{}
	getOpts.ApplyOptions(opts)
	handleGetOptions(span, opts...)

	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)

	if err := k.inner.Get(sctx, key, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to get an object")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
		return err
	}

	span.SetAttributes(
		attribute.String("object.uuid", string(obj.GetUID())),
		attribute.String("object.resourceVersion", obj.GetResourceVersion()),
		attribute.Int64("object.generation", obj.GetGeneration()),
	)

	return nil
}

func handleListOpts(span trace.Span, listOpts *client.ListOptions) {
	if listOpts == nil {
		return
	}
	if listOpts.Namespace != "" {
		span.SetAttributes(attribute.String("listOption.namespace", listOpts.Namespace))
	}
	if fs := listOpts.FieldSelector; fs != nil {
		span.SetAttributes(attribute.String("listOption.fieldSelector", fs.String()))
	}
	if ls := listOpts.LabelSelector; ls != nil {
		span.SetAttributes(attribute.String("listOption.labelSelector", ls.String()))
	}
	if listOpts.Limit != 0 {
		span.SetAttributes(attribute.Int64("listOption.limit", listOpts.Limit))
	}
	if listOpts.Continue != "" {
		span.SetAttributes(attribute.String("listOption.continue", listOpts.Continue))
	}
}

func (k *K8sClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	sctx, span := k.tracer.Start(ctx, "List")
	defer span.End()

	listOpts := &client.ListOptions{}
	for _, opt := range opts {
		opt.ApplyToList(listOpts)
	}
	handleListOpts(span, listOpts)

	k.objectIdentityExtactor.handleClientObjectListAttrs(span, list)

	if err := k.inner.List(sctx, list, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to list objects")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
		return err
	}

	span.AddEvent("successfully listed objects")

	return nil
}

func handleCreateOptions(span trace.Span, opts ...client.CreateOption) {
	if len(opts) == 0 {
		return
	}

	createOpts := &client.CreateOptions{}
	createOpts.ApplyOptions(opts)

	metav1CreateOpts := createOpts.AsCreateOptions()
	if len(metav1CreateOpts.DryRun) > 0 {
		span.SetAttributes(attribute.StringSlice("createOption.dryRun", metav1CreateOpts.DryRun))
	}
	if fm := metav1CreateOpts.FieldManager; fm != "" {
		span.SetAttributes(attribute.String("createOption.fieldManager", fm))
	}
	if fv := metav1CreateOpts.FieldValidation; fv != "" {
		span.SetAttributes(attribute.String("createOption.fieldValidation", fv))
	}
}

func (k *K8sClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	sctx, span := k.tracer.Start(ctx, "Create")
	defer span.End()

	handleCreateOptions(span, opts...)
	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)

	if err := k.inner.Create(sctx, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to create an object")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
		return err
	}
	span.SetAttributes(
		attribute.String("object.uuid", string(obj.GetUID())),
		attribute.String("object.resourceVersion", obj.GetResourceVersion()),
	)

	return nil
}

func handleDeleteOptions(span trace.Span, opts ...client.DeleteOption) {
	if len(opts) == 0 {
		return
	}

	deleteOpts := &client.DeleteOptions{}
	deleteOpts.ApplyOptions(opts)

	if len(deleteOpts.DryRun) > 0 {
		span.SetAttributes(attribute.StringSlice("deleteOption.dryRun", deleteOpts.DryRun))
	}
	if pp := deleteOpts.PropagationPolicy; pp != nil {
		span.SetAttributes(attribute.String("deleteOption.propagationPolicy", string(*pp)))
	}
	if pre := deleteOpts.Preconditions; pre != nil {
		if pre.ResourceVersion != nil {
			span.SetAttributes(attribute.String("deleteOption.preconditions.resourceVersion", *pre.ResourceVersion))
		}
		if pre.UID != nil {
			span.SetAttributes(attribute.String("deleteOption.preconditions.uid", string(*pre.UID)))
		}
	}
	if deleteOpts.GracePeriodSeconds != nil {
		span.SetAttributes(attribute.Int64("deleteOption.gracePeriodSeconds", *deleteOpts.GracePeriodSeconds))
	}
}

func (k *K8sClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	sctx, span := k.tracer.Start(ctx, "Delete", trace.WithAttributes(attribute.String("object.name", obj.GetName())))
	defer span.End()

	if ns := obj.GetNamespace(); ns != "" {
		span.SetAttributes(attribute.String("object.namespace", ns))
	}
	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)
	handleDeleteOptions(span, opts...)

	if err := k.inner.Delete(sctx, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to delete an object")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
		return err
	}

	return nil
}

func handleUpdateOpts(span trace.Span, opts *metav1.UpdateOptions) {
	if opts == nil {
		return
	}
	if len(opts.DryRun) > 0 {
		span.SetAttributes(attribute.StringSlice("updateOption.dryRun", opts.DryRun))
	}
	if fm := opts.FieldManager; fm != "" {
		span.SetAttributes(attribute.String("updateOption.fieldManager", fm))
	}
	if fv := opts.FieldValidation; fv != "" {
		span.SetAttributes(attribute.String("updateOption.fieldValidation", fv))
	}
}

func (k *K8sClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	sctx, span := k.tracer.Start(ctx, "Update")
	defer span.End()

	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)

	updateOpts := &client.UpdateOptions{}
	updateOpts.ApplyOptions(opts)
	handleUpdateOpts(span, updateOpts.AsUpdateOptions())

	if err := k.inner.Update(sctx, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to update an object")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
		return err
	}

	return nil
}

func handlePatchOpts(span trace.Span, opts *metav1.PatchOptions) {
	if opts == nil {
		return
	}
	if len(opts.DryRun) > 0 {
		span.SetAttributes(attribute.StringSlice("patchOption.dryRun", opts.DryRun))
	}
	if fm := opts.FieldManager; fm != "" {
		span.SetAttributes(attribute.String("patchOption.fieldManager", fm))
	}
	if fv := opts.FieldValidation; fv != "" {
		span.SetAttributes(attribute.String("patchOption.fieldValidation", fv))
	}
	if opts.Force != nil {
		span.SetAttributes(attribute.Bool("patchOption.Force", *opts.Force))
	}
}

func (k *K8sClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	sctx, span := k.tracer.Start(ctx, "Patch", trace.WithAttributes(attribute.String("object.name", obj.GetName())))
	defer span.End()

	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)

	span.SetAttributes(attribute.String("patch.type", string(patch.Type())))

	patchOpts := &client.PatchOptions{}
	patchOpts.ApplyOptions(opts)
	handlePatchOpts(span, patchOpts.AsPatchOptions())

	if err := k.inner.Patch(sctx, obj, patch, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to patch an object")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
		return err
	}

	return nil
}

func (k *K8sClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	sctx, span := k.tracer.Start(ctx, "DeleteAllOf")
	defer span.End()

	k.objectIdentityExtactor.handleClientObjectAttrs(span, false, obj)

	delOpts := &client.DeleteAllOfOptions{}
	delOpts.ApplyOptions(opts)

	handleListOpts(span, &delOpts.ListOptions)
	handleDeleteOptions(span, &delOpts.DeleteOptions)

	if err := k.inner.DeleteAllOf(sctx, obj, opts...); err != nil {
		reason := apierrors.ReasonForError(err)
		span.AddEvent("failed to run DeleteAllOf")
		span.SetAttributes(attribute.String("reasonForError", string(reason)))
		SetSpanErr(span, err)
		return err
	}

	return nil
}

func (k *K8sClient) RESTMapper() meta.RESTMapper {
	return k.inner.RESTMapper()
}

func (k *K8sClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	return k.inner.GroupVersionKindFor(obj)
}

func (k *K8sClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	return k.inner.IsObjectNamespaced(obj)
}

func (k *K8sClient) Scheme() *runtime.Scheme {
	return k.inner.Scheme()
}

func (k *K8sClient) Status() client.SubResourceWriter {
	return &K8sSubresourceClient{
		inner:                 k.SubResource("status"),
		groupVersionKindForFn: k.GroupVersionKindFor,
		tracer:                k.traceProvider.Tracer("KubernetesClient.subresource.status"),
	}
}

func (k *K8sClient) SubResource(subResource string) client.SubResourceClient {
	return &K8sSubresourceClient{
		inner:                 k.inner.SubResource(subResource),
		groupVersionKindForFn: k.GroupVersionKindFor,
		tracer:                k.traceProvider.Tracer(fmt.Sprintf("KubernetesClient.subresource.%s", subResource)),
	}
}
