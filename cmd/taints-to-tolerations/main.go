package main

import (
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"
)

func main() {
	taints := []corev1.Taint{}
	if err := utilyaml.NewYAMLOrJSONDecoder(os.Stdin, 2048).Decode(&taints); err != nil {
		panic(err)
	}

	tolerations, err := yaml.Marshal(convert(taints))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tolerations))
}

func convert(in []corev1.Taint) []corev1.Toleration {
	out := make([]corev1.Toleration, len(in))
	for i, taint := range in {
		tol := corev1.Toleration{
			Key:      taint.Key,
			Operator: "",
			Value:    "",
			Effect:   taint.Effect,
		}

		if taint.Value != "" {
			tol.Value = taint.Value
			tol.Operator = corev1.TolerationOpEqual
		} else {
			tol.Operator = corev1.TolerationOpExists
		}

		out[i] = tol
	}
	return out
}
