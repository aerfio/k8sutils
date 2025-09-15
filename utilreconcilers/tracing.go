package utilreconcilers

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"aerf.io/k8sutils/k8stracing"
)

type WithTracingReconciler struct {
	inner  reconcile.Reconciler
	tracer trace.Tracer
}

func NewWithTracingReconciler(inner reconcile.Reconciler, tracer trace.Tracer) *WithTracingReconciler {
	return &WithTracingReconciler{
		inner:  inner,
		tracer: tracer,
	}
}

func (r *WithTracingReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	sctx, span := r.tracer.Start(ctx, "Reconcile",
		trace.WithAttributes(
			attribute.String("request.name", req.Name),
			attribute.String("request.namespace", req.Namespace),
			attribute.String("reconcile.id", string(controller.ReconcileIDFromContext(ctx))),
		),
	)
	defer span.End()

	if spanCtx := span.SpanContext(); spanCtx.IsValid() && span.IsRecording() {
		sctx = ctrl.LoggerInto(sctx,
			ctrl.LoggerFrom(sctx, "trace-id", spanCtx.TraceID().String()),
		)
	}

	res, err := r.inner.Reconcile(sctx, req)
	if err != nil {
		span.SetAttributes(attribute.String("errReason", string(apierrors.ReasonForError(err))))
		k8stracing.SetSpanErr(span, err)
		return res, err
	}

	span.SetAttributes(attribute.Bool("requeue", res.Requeue), attribute.String("requeueAfter", res.RequeueAfter.String())) //nolint:staticcheck // I know Requeue is deprecated
	return res, err
}
