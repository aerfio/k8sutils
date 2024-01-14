package utilreconcilers

import (
	"context"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func RequeueOnConflict(r reconcile.Reconciler) reconcile.Reconciler {
	return &requeueOnConflict{internal: r}
}

type requeueOnConflict struct {
	internal reconcile.Reconciler
}

func (r *requeueOnConflict) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	result, err := r.internal.Reconcile(ctx, req)
	if apierrors.IsConflict(err) {
		return reconcile.Result{Requeue: true}, nil
	}
	return result, err
}
