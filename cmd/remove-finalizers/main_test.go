package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Test_readObjects(t *testing.T) {
	tests := []struct {
		name    string
		object  string
		want    []*unstructured.Unstructured
		wantErr bool
	}{
		{
			name: "simple case with 1 item",
			object: `kind: ConfigMap
apiVersion: v1
metadata:
  name: test`,
			want: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"kind":       "ConfigMap",
						"apiVersion": "v1",
						"metadata": map[string]any{
							"name": "test",
						},
					},
				},
			},
		},
		{
			name: "items split with ---",
			object: `kind: ConfigMap
apiVersion: v1
metadata:
  name: test
---
kind: Secret
apiVersion: v1
metadata:
  name: test-2`,

			want: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"kind":       "ConfigMap",
						"apiVersion": "v1",
						"metadata": map[string]any{
							"name": "test",
						},
					},
				},
				{
					Object: map[string]any{
						"kind":       "Secret",
						"apiVersion": "v1",
						"metadata": map[string]any{
							"name": "test-2",
						},
					},
				},
			},
		},
		{
			name: "items list",
			object: `kind: List
apiVersion: v1
items: 
- kind: ConfigMap
  apiVersion: v1
  metadata:
    name: test
- kind: Secret
  apiVersion: v1
  metadata:
    name: test-2`,

			want: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"kind":       "ConfigMap",
						"apiVersion": "v1",
						"metadata": map[string]any{
							"name": "test",
						},
					},
				},
				{
					Object: map[string]any{
						"kind":       "Secret",
						"apiVersion": "v1",
						"metadata": map[string]any{
							"name": "test-2",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := strings.NewReader(tt.object)
			got, err := readObjects(obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("readObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("readObjects() got = %v, want %v, diff:\n%s", got, tt.want, diff)
			}
		})
	}
}
