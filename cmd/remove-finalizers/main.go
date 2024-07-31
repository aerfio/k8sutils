package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	objs, err := readObjects(os.Stdin)
	if err != nil {
		panic(err)
	}

	restCfg := config.GetConfigOrDie()
	cli, err := client.New(restCfg, client.Options{})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, obj := range objs {
		accessor := obj.GetName()
		if ns := obj.GetNamespace(); ns != "" {
			accessor = fmt.Sprintf("%s/%s", obj.GetName(), ns)
		}
		obj.SetFinalizers([]string{})

		fmt.Printf("removing finalizers from %s %s\n", obj.GetObjectKind().GroupVersionKind().Kind, accessor)
		if err := cli.Update(ctx, obj); err != nil {
			panic(err)
		}
	}
}

// copied from github.com/fluxcd/ssa/utils - thanks!
// ReadObjects decodes the YAML or JSON documents from the given reader into unstructured Kubernetes API objects.
// The documents which do not subscribe to the Kubernetes Object interface, are silently dropped from the result.
func readObjects(r io.Reader) ([]*unstructured.Unstructured, error) {
	reader := utilyaml.NewYAMLOrJSONDecoder(r, 2048)
	objects := make([]*unstructured.Unstructured, 0)

	for {
		obj := &unstructured.Unstructured{}
		err := reader.Decode(obj)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			return objects, err
		}

		if obj.IsList() {
			err = obj.EachListItem(func(item runtime.Object) error {
				obj := item.(*unstructured.Unstructured) //nolint:errcheck,forcetypeassert // item must be *unstructured.unstructured
				objects = append(objects, obj)
				return nil
			})
			if err != nil {
				return objects, err
			}
			continue
		}
	}

	return objects, nil
}
