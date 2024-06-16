package main

import (
	"context"
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"aerf.io/k8sutils/config"
)

func main() {
	obj := &unstructured.Unstructured{}
	if err := yaml.NewYAMLOrJSONDecoder(os.Stdin, 2048).Decode(obj); err != nil {
		panic(err)
	}
	obj.SetFinalizers([]string{})
	restCfg := config.GetConfigOrDie()
	cli, err := client.New(restCfg, client.Options{})
	if err != nil {
		panic(err)
	}
	if err := cli.Update(context.Background(), obj); err != nil {
		panic(err)
	}
}
