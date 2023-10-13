package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	coordinationv1 "k8s.io/api/coordination/v1"
	coordinationv1beta1 "k8s.io/api/coordination/v1beta1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	eventsv1 "k8s.io/api/events/v1"
	networkingv1 "k8s.io/api/networking/v1"
	policyv1 "k8s.io/api/policy/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	schedulingv1 "k8s.io/api/scheduling/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type versionedPackageFns struct {
	gv            schema.GroupVersion
	ignoredKinds  []string
	addToSchemeFn func(*runtime.Scheme) error
}

func run() error {
	for _, input := range []struct {
		pckgRoot            string
		versionedPackageFns []versionedPackageFns
	}{
		{
			pckgRoot: "core",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            corev1.SchemeGroupVersion,
					ignoredKinds:  []string{"Status", "List", "ComponentStatus", "ComponentStatusList", "PodStatusResult"},
					addToSchemeFn: corev1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "rbac",
			versionedPackageFns: []versionedPackageFns{
				{
					gv: rbacv1.SchemeGroupVersion,
					// ignoredKinds:  []string{"Status", "List", "ComponentStatus", "ComponentStatusList", "PodStatusResult"},
					addToSchemeFn: rbacv1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "apps",
			versionedPackageFns: []versionedPackageFns{
				{
					gv: appsv1.SchemeGroupVersion,
					// ignoredKinds:  []string{"Status", "List", "ComponentStatus", "ComponentStatusList", "PodStatusResult"},
					addToSchemeFn: appsv1.AddToScheme,
				}, {
					gv:            appsv1beta1.SchemeGroupVersion,
					addToSchemeFn: appsv1beta1.AddToScheme,
				}, {
					gv:            appsv1beta2.SchemeGroupVersion,
					addToSchemeFn: appsv1beta2.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "batch",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            batchv1.SchemeGroupVersion,
					addToSchemeFn: batchv1.AddToScheme,
				}, {
					gv:            batchv1beta1.SchemeGroupVersion,
					addToSchemeFn: batchv1beta1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "autoscaling",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            autoscalingv1.SchemeGroupVersion,
					addToSchemeFn: autoscalingv1.AddToScheme,
				},
				{
					gv:            autoscalingv2.SchemeGroupVersion,
					addToSchemeFn: autoscalingv2.AddToScheme,
				},
				{
					gv:            autoscalingv2beta1.SchemeGroupVersion,
					addToSchemeFn: autoscalingv2beta1.AddToScheme,
				},
				{
					gv:            autoscalingv2beta2.SchemeGroupVersion,
					addToSchemeFn: autoscalingv2beta2.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "coordination",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            coordinationv1.SchemeGroupVersion,
					addToSchemeFn: coordinationv1.AddToScheme,
				},
				{
					gv:            coordinationv1beta1.SchemeGroupVersion,
					addToSchemeFn: coordinationv1beta1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "events",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            eventsv1.SchemeGroupVersion,
					addToSchemeFn: eventsv1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "networking",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            networkingv1.SchemeGroupVersion,
					addToSchemeFn: networkingv1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "policy",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            policyv1.SchemeGroupVersion,
					addToSchemeFn: policyv1.AddToScheme,
					ignoredKinds:  []string{"Eviction"},
				},
			},
		},
		{
			pckgRoot: "discovery",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            discoveryv1.SchemeGroupVersion,
					addToSchemeFn: discoveryv1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "scheduling",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            schedulingv1.SchemeGroupVersion,
					addToSchemeFn: schedulingv1.AddToScheme,
				},
			},
		},
		{
			pckgRoot: "storage",
			versionedPackageFns: []versionedPackageFns{
				{
					gv:            storagev1.SchemeGroupVersion,
					addToSchemeFn: storagev1.AddToScheme,
				},
			},
		},
	} {
		for _, elem := range input.versionedPackageFns {
			pckg := fmt.Sprintf("%s%sgvk", input.pckgRoot, elem.gv.Version)
			outFile := fmt.Sprintf("./%s/%s.go", pckg, pckg)
			if err := createPackage(pckg, fmt.Sprintf("k8s.io/api/%s/%s", input.pckgRoot, elem.gv.Version), elem.gv, elem.ignoredKinds, outFile, elem.addToSchemeFn); err != nil {
				return err
			}
		}
	}

	return nil
}

type K8sObject interface {
	runtime.Object
	metav1.Object
}
type K8sListObject interface {
	runtime.Object
	metav1.ListInterface
}

func groupObjAndList(knownTypes map[string]reflect.Type) []string {
	out := make([]string, 0, len(knownTypes))
	lists := []string{}
	nonLists := []string{}
	for kind := range knownTypes {
		if strings.HasSuffix(kind, "List") {
			lists = append(lists, kind)
		} else {
			nonLists = append(nonLists, kind)
		}
	}
	sort.Strings(nonLists)
	for _, elem := range nonLists {
		out = append(out, elem)
		if slices.Contains(lists, elem+"List") {
			out = append(out, elem+"List")
			idx := slices.Index(lists, elem+"List")
			lists = append(lists[:idx], lists[idx+1:]...)
		}
	}
	out = append(out, lists...)
	return out
}

func createPackage(pckg, k8sPackage string, gv schema.GroupVersion, ignoreKinds []string, outFilePath string, addToSchemeFn func(*runtime.Scheme) error) error {
	scheme := runtime.NewScheme()
	if err := addToSchemeFn(scheme); err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "package %s\n\n", pckg)
	fmt.Fprintf(buf, "import (\n\t\"reflect\"\n\t%q\n)\n", k8sPackage)

	for _, kind := range groupObjAndList(scheme.KnownTypes(gv)) {
		if slices.Contains(ignoreKinds, kind) {
			continue
		}
		obj, err := scheme.New(gv.WithKind(kind))
		if err != nil {
			return err
		}
		_, isK8sObj := obj.(K8sObject)
		_, isK8sObjList := obj.(K8sListObject)
		if isK8sObj || isK8sObjList {
			fmt.Fprintf(buf, "var %s = %s.SchemeGroupVersion.WithKind(reflect.TypeOf(&%s.%s{}).Name())\n", kind, filepath.Base(k8sPackage), filepath.Base(k8sPackage), kind)
		}
	}

	if err := os.Mkdir("./"+filepath.Dir(outFilePath), os.FileMode(0o755)); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}

	file, err := os.OpenFile(outFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, buf); err != nil {
		return err
	}

	return nil
}
