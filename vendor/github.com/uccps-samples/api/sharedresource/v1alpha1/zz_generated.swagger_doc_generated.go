package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_SharedConfigMap = map[string]string{
	"":       "SharedConfigMap allows a ConfigMap to be shared across namespaces. Pods can mount the shared ConfigMap by adding a CSI volume to the pod specification using the \"csi.sharedresource.uccp.io\" CSI driver and a reference to the SharedConfigMap in the volume attributes:\n\nspec:\n volumes:\n - name: shared-configmap\n   csi:\n     driver: csi.sharedresource.uccp.io\n     volumeAttributes:\n       sharedConfigMap: my-share\n\nFor the mount to be successful, the pod's service account must be granted permission to 'use' the named SharedConfigMap object within its namespace with an appropriate Role and RoleBinding. For compactness, here are example `oc` invocations for creating such Role and RoleBinding objects.\n\n `oc create role shared-resource-my-share --verb=use --resource=sharedconfigmaps.sharedresource.uccp.io --resource-name=my-share`\n `oc create rolebinding shared-resource-my-share --role=shared-resource-my-share --serviceaccount=my-namespace:default`\n\nShared resource objects, in this case ConfigMaps, have default permissions of list, get, and watch for system authenticated users.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.",
	"spec":   "spec is the specification of the desired shared configmap",
	"status": "status is the observed status of the shared configmap",
}

func (SharedConfigMap) SwaggerDoc() map[string]string {
	return map_SharedConfigMap
}

var map_SharedConfigMapList = map[string]string{
	"": "SharedConfigMapList contains a list of SharedConfigMap objects.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.",
}

func (SharedConfigMapList) SwaggerDoc() map[string]string {
	return map_SharedConfigMapList
}

var map_SharedConfigMapReference = map[string]string{
	"":          "SharedConfigMapReference contains information about which ConfigMap to share",
	"name":      "name represents the name of the ConfigMap that is being referenced.",
	"namespace": "namespace represents the namespace where the referenced ConfigMap is located.",
}

func (SharedConfigMapReference) SwaggerDoc() map[string]string {
	return map_SharedConfigMapReference
}

var map_SharedConfigMapSpec = map[string]string{
	"":             "SharedConfigMapSpec defines the desired state of a SharedConfigMap",
	"configMapRef": "configMapRef is a reference to the ConfigMap to share",
	"description":  "description is a user readable explanation of what the backing resource provides.",
}

func (SharedConfigMapSpec) SwaggerDoc() map[string]string {
	return map_SharedConfigMapSpec
}

var map_SharedConfigMapStatus = map[string]string{
	"":           "SharedSecretStatus contains the observed status of the shared resource",
	"conditions": "conditions represents any observations made on this particular shared resource by the underlying CSI driver or Share controller.",
}

func (SharedConfigMapStatus) SwaggerDoc() map[string]string {
	return map_SharedConfigMapStatus
}

var map_SharedSecret = map[string]string{
	"":       "SharedSecret allows a Secret to be shared across namespaces. Pods can mount the shared Secret by adding a CSI volume to the pod specification using the \"csi.sharedresource.uccp.io\" CSI driver and a reference to the SharedSecret in the volume attributes:\n\nspec:\n volumes:\n - name: shared-secret\n   csi:\n     driver: csi.sharedresource.uccp.io\n     volumeAttributes:\n       sharedSecret: my-share\n\nFor the mount to be successful, the pod's service account must be granted permission to 'use' the named SharedSecret object within its namespace with an appropriate Role and RoleBinding. For compactness, here are example `oc` invocations for creating such Role and RoleBinding objects.\n\n `oc create role shared-resource-my-share --verb=use --resource=sharedsecrets.sharedresource.uccp.io --resource-name=my-share`\n `oc create rolebinding shared-resource-my-share --role=shared-resource-my-share --serviceaccount=my-namespace:default`\n\nShared resource objects, in this case Secrets, have default permissions of list, get, and watch for system authenticated users.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.",
	"spec":   "spec is the specification of the desired shared secret",
	"status": "status is the observed status of the shared secret",
}

func (SharedSecret) SwaggerDoc() map[string]string {
	return map_SharedSecret
}

var map_SharedSecretList = map[string]string{
	"": "SharedSecretList contains a list of SharedSecret objects.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.",
}

func (SharedSecretList) SwaggerDoc() map[string]string {
	return map_SharedSecretList
}

var map_SharedSecretReference = map[string]string{
	"":          "SharedSecretReference contains information about which Secret to share",
	"name":      "name represents the name of the Secret that is being referenced.",
	"namespace": "namespace represents the namespace where the referenced Secret is located.",
}

func (SharedSecretReference) SwaggerDoc() map[string]string {
	return map_SharedSecretReference
}

var map_SharedSecretSpec = map[string]string{
	"":            "SharedSecretSpec defines the desired state of a SharedSecret",
	"secretRef":   "secretRef is a reference to the Secret to share",
	"description": "description is a user readable explanation of what the backing resource provides.",
}

func (SharedSecretSpec) SwaggerDoc() map[string]string {
	return map_SharedSecretSpec
}

var map_SharedSecretStatus = map[string]string{
	"":           "SharedSecretStatus contains the observed status of the shared resource",
	"conditions": "conditions represents any observations made on this particular shared resource by the underlying CSI driver or Share controller.",
}

func (SharedSecretStatus) SwaggerDoc() map[string]string {
	return map_SharedSecretStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
