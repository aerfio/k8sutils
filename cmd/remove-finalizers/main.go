package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"aerf.io/k8sutils/config"
)

func main() {
	obj := []*unstructured.Unstructured{}

	if err := yaml.NewYAMLOrJSONDecoder(os.Stdin, 2048).Decode(obj); err != nil {
		panic(err)
	}
	objs, err := readObjects(os.Stdin)
	if err != nil {
		panic(err)
	}

	for i := range objs {
		objs[i].SetFinalizers([]string{})
	}

	restCfg := config.GetConfigOrDie()
	cli, err := client.New(restCfg, client.Options{})
	if err != nil {
		panic(err)
	}

	for _, obj := range objs {
		accessor := obj.GetName()
		if ns := obj.GetNamespace(); ns != "" {
			accessor = fmt.Sprintf("%s/%s", obj.GetName(), ns)
		}
		fmt.Printf("removing finalizers from %s %s", obj.GetObjectKind().GroupVersionKind().Kind, accessor)
		if err := cli.Update(context.Background(), obj); err != nil {
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
				obj := item.(*unstructured.Unstructured)
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
