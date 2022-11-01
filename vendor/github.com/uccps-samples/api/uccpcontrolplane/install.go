package uccpcontrolplane

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	uccpcontrolplanev1 "github.com/uccps-samples/api/uccpcontrolplane/v1"
)

const (
	GroupName = "uccpcontrolplane.config.uccp.io"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(uccpcontrolplanev1.Install)
	// Install is a function which adds every version of this group to a scheme
	Install = schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return schema.GroupResource{Group: GroupName, Resource: resource}
}

func Kind(kind string) schema.GroupKind {
	return schema.GroupKind{Group: GroupName, Kind: kind}
}
