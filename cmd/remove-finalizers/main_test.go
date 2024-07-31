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
  name: xd`,
			want: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"kind":       "ConfigMap",
						"apiVersion": "v1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.object)
			got, err := readObjects(strings.NewReader(tt.object))
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
