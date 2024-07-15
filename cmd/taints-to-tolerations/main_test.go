package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
)

func Test_convert(t *testing.T) {
	tests := []struct {
		name        string
		taints      string
		tolerations string
	}{
		{
			name: "1 taint with value",
			taints: `
- effect: NoSchedule
  key: some-key
  value: "true"
`,
			tolerations: `
- effect: NoSchedule
  key: some-key
  operator: Equal
  value: "true"
`,
		},
		{
			name: "1 taint with no value",
			taints: `
- effect: NoSchedule
  key: some-key
`,
			tolerations: `
- effect: NoSchedule
  key: some-key
  operator: Exists
`,
		},
		{
			name: "multiple taints",
			taints: `
- effect: NoSchedule
  key: some-key
- effect: NoSchedule
  key: some-other-key
  value: "true"
`,
			tolerations: `
- effect: NoSchedule
  key: some-key
  operator: Exists
- effect: NoSchedule
  key: some-other-key
  operator: Equal
  value: "true"
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taints := mustReadObject[corev1.Taint](t, tt.taints)
			wantTolerations := mustReadObject[corev1.Toleration](t, tt.tolerations)

			got := convert(taints)
			if diff := cmp.Diff(got, wantTolerations); diff != "" {
				t.Errorf("convert() = %v, want %v diff:\n%s", got, wantTolerations, diff)
			}
		})
	}
}

func mustReadObject[T any](t *testing.T, objRaw string) []T {
	reader := utilyaml.NewYAMLOrJSONDecoder(strings.NewReader(objRaw), 2048)
	out := make([]T, 0)

	if err := reader.Decode(&out); err != nil {
		t.Fatalf("failed to decode input obj into target slice: %s", err)
	}

	return out
}
