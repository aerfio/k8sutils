package cachedebug

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// HIGHLY EXPERIMENTAL!
func NewDebugCacheTransformer(httpSrvAddr string, scheme *runtime.Scheme) *DebugTransform {
	return &DebugTransform{
		cache:       make(map[schema.GroupVersionKind][]string),
		mu:          new(sync.Mutex),
		httpSrvAddr: httpSrvAddr,
		scheme:      scheme,
	}
}

type DebugTransform struct {
	cache       map[schema.GroupVersionKind][]string
	mu          *sync.Mutex
	httpSrvAddr string
	scheme      *runtime.Scheme
}

func (dt *DebugTransform) NeedLeaderElection() bool {
	// caches are filled in only in leader
	return true
}

func (dt *DebugTransform) getCache() map[schema.GroupVersionKind][]string {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	return dt.cache
}

func (dt *DebugTransform) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/cachedebug", func(w http.ResponseWriter, r *http.Request) {
		marshallable := make(map[string][]string)
		for key, val := range dt.getCache() {
			marshallable[key.String()] = val
		}
		out, err := json.Marshal(marshallable)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal internal debug cache content: %s", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out) // nolint:errcheck
	})
	srv := http.Server{ // nolint:gosec // this shouldnt be exposed so lint err doesnt matter
		Addr:    dt.httpSrvAddr,
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		sctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(sctx) // nolint:errcheck
	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

var (
	_ manager.Runnable               = &DebugTransform{}
	_ manager.LeaderElectionRunnable = &DebugTransform{}
)

func (dt *DebugTransform) TransformFn(fns ...toolscache.TransformFunc) toolscache.TransformFunc {
	return func(obj any) (any, error) {
		var err error
		for i := range fns {
			obj, err = fns[i](obj)
			if err != nil {
				return obj, err
			}
		}

		if robj, ok := obj.(client.Object); ok {

			gvk := robj.GetObjectKind().GroupVersionKind()
			if gvk == (schema.GroupVersionKind{}) {
				gvk, err = apiutil.GVKForObject(robj, dt.scheme)
				if err != nil {
					return nil, err
				}
			}

			dt.mu.Lock()
			defer dt.mu.Unlock()
			nnt := dt.cache[gvk]
			typeOfObj := reflect.TypeOf(robj)
			objKey := typeOfObj.String()
			if !slices.Contains(nnt, objKey) {
				nnt = append(nnt, objKey)
			}
			dt.cache[gvk] = nnt
		}

		return obj, nil
	}
}
