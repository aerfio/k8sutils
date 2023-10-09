package rbacv1gvk

import (
	"reflect"

	"k8s.io/api/rbac/v1"
)

var (
	RoleBinding            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.RoleBinding{}).Name())
	RoleBindingList        = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.RoleBindingList{}).Name())
	ClusterRoleBinding     = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ClusterRoleBinding{}).Name())
	ClusterRoleBindingList = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ClusterRoleBindingList{}).Name())
	ClusterRole            = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ClusterRole{}).Name())
	ClusterRoleList        = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.ClusterRoleList{}).Name())
	Role                   = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.Role{}).Name())
	RoleList               = v1.SchemeGroupVersion.WithKind(reflect.TypeOf(&v1.RoleList{}).Name())
)
