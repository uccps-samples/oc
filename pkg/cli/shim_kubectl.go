package cli

import (
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/openshiftpatch"
	kclientcmd "k8s.io/client-go/tools/clientcmd"
	kclientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog/v2"
	kcmdset "k8s.io/kubectl/pkg/cmd/set"
	describeversioned "k8s.io/kubectl/pkg/describe"
	"k8s.io/kubectl/pkg/generate/versioned"
	"k8s.io/kubectl/pkg/polymorphichelpers"

	"github.com/uccps-samples/library-go/pkg/legacyapi/legacygroupification"
	"github.com/uccps-samples/oc/pkg/helpers/clientcmd"
	"github.com/uccps-samples/oc/pkg/helpers/originpolymorphichelpers"
)

func shimKubectlForOc() {
	// we only need this change for `oc`.  `kubectl` should behave as close to `kubectl` as we can
	// if we call this factory construction method, we want the openshift style config loading
	kclientcmd.ErrEmptyConfig = genericclioptions.ErrEmptyConfig
	kclientcmd.ClusterDefaults = kclientcmdapi.Cluster{Server: os.Getenv("KUBERNETES_MASTER")}
	kcmdset.ParseDockerImageReferenceToStringFunc = clientcmd.ParseDockerImageReferenceToStringFunc
	openshiftpatch.OAPIToGroupified = oAPIToGroupified
	openshiftpatch.OAPIToGroupifiedGVK = oAPIToGroupifiedGVK
	openshiftpatch.IsOAPIFn = legacygroupification.IsOAPI
	openshiftpatch.IsOC = true
	kclientcmd.UseModifyConfigLock = false

	// update polymorphic helpers
	polymorphichelpers.AttachablePodForObjectFn = originpolymorphichelpers.NewAttachablePodForObjectFn(polymorphichelpers.AttachablePodForObjectFn)
	polymorphichelpers.CanBeExposedFn = originpolymorphichelpers.NewCanBeExposedFn(polymorphichelpers.CanBeExposedFn)
	polymorphichelpers.HistoryViewerFn = originpolymorphichelpers.NewHistoryViewerFn(polymorphichelpers.HistoryViewerFn)
	polymorphichelpers.LogsForObjectFn = originpolymorphichelpers.NewLogsForObjectFn(polymorphichelpers.LogsForObjectFn)
	polymorphichelpers.MapBasedSelectorForObjectFn = originpolymorphichelpers.NewMapBasedSelectorForObjectFn(polymorphichelpers.MapBasedSelectorForObjectFn)
	polymorphichelpers.ObjectPauserFn = originpolymorphichelpers.NewObjectPauserFn(polymorphichelpers.ObjectPauserFn)
	polymorphichelpers.ObjectResumerFn = originpolymorphichelpers.NewObjectResumerFn(polymorphichelpers.ObjectResumerFn)
	polymorphichelpers.PortsForObjectFn = originpolymorphichelpers.NewPortsForObjectFn(polymorphichelpers.PortsForObjectFn)
	polymorphichelpers.ProtocolsForObjectFn = originpolymorphichelpers.NewProtocolsForObjectFn(polymorphichelpers.ProtocolsForObjectFn)
	polymorphichelpers.RollbackerFn = originpolymorphichelpers.NewRollbackerFn(polymorphichelpers.RollbackerFn)
	polymorphichelpers.StatusViewerFn = originpolymorphichelpers.NewStatusViewerFn(polymorphichelpers.StatusViewerFn)
	polymorphichelpers.UpdatePodSpecForObjectFn = originpolymorphichelpers.NewUpdatePodSpecForObjectFn(polymorphichelpers.UpdatePodSpecForObjectFn)

	// update some functions we inject
	versioned.GeneratorFn = originpolymorphichelpers.NewGeneratorsFn(versioned.GeneratorFn)
	describeversioned.DescriberFn = originpolymorphichelpers.NewDescriberFn(describeversioned.DescriberFn)
}

func oAPIToGroupified(uncast runtime.Object, gvk *schema.GroupVersionKind) {
	before := gvk.Group
	legacygroupification.OAPIToGroupified(uncast, gvk)
	if before != gvk.Group {
		klog.Warningf("Using non-groupfied API resources is deprecated and will be removed in a future release, update apiVersion to %q for your resource", gvk.GroupVersion())
	}
}

func oAPIToGroupifiedGVK(gvk *schema.GroupVersionKind) {
	before := gvk.Group
	legacygroupification.OAPIToGroupifiedGVK(gvk)
	if before != gvk.Group {
		klog.Warningf("Using non-groupfied API resources is deprecated and will be removed in a future release, update apiVersion to %q for your resource", gvk.GroupVersion())
	}
}
