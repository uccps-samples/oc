package main

import (
	goflag "flag"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/pflag"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachineryruntime "k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"k8s.io/kubectl/pkg/scheme"

	"github.com/uccps-samples/api/apps"
	"github.com/uccps-samples/api/authorization"
	"github.com/uccps-samples/api/build"
	"github.com/uccps-samples/api/image"
	"github.com/uccps-samples/api/network"
	"github.com/uccps-samples/api/oauth"
	"github.com/uccps-samples/api/project"
	quotav1 "github.com/uccps-samples/api/quota/v1"
	"github.com/uccps-samples/api/route"
	securityv1 "github.com/uccps-samples/api/security/v1"
	"github.com/uccps-samples/api/template"
	"github.com/uccps-samples/api/user"
	"github.com/uccps-samples/library-go/pkg/serviceability"

	"github.com/uccps-samples/oc/pkg/cli"
	"github.com/uccps-samples/oc/pkg/helpers/legacy"
	"github.com/uccps-samples/oc/pkg/version"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func injectLoglevelFlag(flags *pflag.FlagSet) {
	from := goflag.CommandLine
	if flag := from.Lookup("v"); flag != nil {
		level := flag.Value.(*klog.Level)
		levelPtr := (*int32)(level)
		flags.Int32Var(levelPtr, "loglevel", 0, "Set the level of log output (0-10)")
		if flags.Lookup("v") == nil {
			flags.Int32Var(levelPtr, "v", 0, "Set the level of log output (0-10)")
		}
	}
}

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()
	defer serviceability.BehaviorOnPanic(os.Getenv("UCCP_ON_PANIC"), version.Get())()
	defer serviceability.Profile(os.Getenv("UCCP_PROFILE")).Stop()

	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// Prevents race condition present in vendored version of Docker.
	// See: https://github.com/moby/moby/issues/39859
	os.Setenv("MOBY_DISABLE_PIGZ", "true")

	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	injectLoglevelFlag(pflag.CommandLine)

	// the kubectl scheme expects to have all the recognizable external types it needs to consume.  Install those here.
	// We can't use the "normal" scheme because apply will use that to build stategic merge patches on CustomResources
	utilruntime.Must(apps.Install(scheme.Scheme))
	utilruntime.Must(authorization.Install(scheme.Scheme))
	utilruntime.Must(build.Install(scheme.Scheme))
	utilruntime.Must(image.Install(scheme.Scheme))
	utilruntime.Must(network.Install(scheme.Scheme))
	utilruntime.Must(oauth.Install(scheme.Scheme))
	utilruntime.Must(project.Install(scheme.Scheme))
	utilruntime.Must(installNonCRDQuota(scheme.Scheme))
	utilruntime.Must(route.Install(scheme.Scheme))
	utilruntime.Must(installNonCRDSecurity(scheme.Scheme))
	utilruntime.Must(template.Install(scheme.Scheme))
	utilruntime.Must(user.Install(scheme.Scheme))
	legacy.InstallExternalLegacyAll(scheme.Scheme)

	basename := filepath.Base(os.Args[0])
	command := cli.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func installNonCRDSecurity(scheme *apimachineryruntime.Scheme) error {
	scheme.AddKnownTypes(securityv1.GroupVersion,
		&securityv1.PodSecurityPolicySubjectReview{},
		&securityv1.PodSecurityPolicySelfSubjectReview{},
		&securityv1.PodSecurityPolicyReview{},
		&securityv1.RangeAllocation{},
		&securityv1.RangeAllocationList{},
	)
	if err := corev1.AddToScheme(scheme); err != nil {
		return err
	}
	metav1.AddToGroupVersion(scheme, securityv1.GroupVersion)
	return nil
}

func installNonCRDQuota(scheme *apimachineryruntime.Scheme) error {
	scheme.AddKnownTypes(securityv1.GroupVersion,
		&quotav1.AppliedClusterResourceQuota{},
		&quotav1.AppliedClusterResourceQuotaList{},
	)
	if err := corev1.AddToScheme(scheme); err != nil {
		return err
	}
	metav1.AddToGroupVersion(scheme, quotav1.GroupVersion)
	return nil
}
