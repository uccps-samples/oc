package newapp

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/uccps-samples/api"
	"github.com/uccps-samples/api/apps"
	"github.com/uccps-samples/api/authorization"
	"github.com/uccps-samples/api/build"
	"github.com/uccps-samples/api/image"
	"github.com/uccps-samples/api/network"
	"github.com/uccps-samples/api/oauth"
	"github.com/uccps-samples/api/project"
	"github.com/uccps-samples/api/quota"
	"github.com/uccps-samples/api/route"
	"github.com/uccps-samples/api/security"
	"github.com/uccps-samples/api/template"
	"github.com/uccps-samples/api/user"
)

// we need a scheme that contains ONLY groupped APIs for newapp
var newAppScheme = runtime.NewScheme()

func init() {
	utilruntime.Must(api.InstallKube(newAppScheme))

	utilruntime.Must(apps.Install(newAppScheme))
	utilruntime.Must(authorization.Install(newAppScheme))
	utilruntime.Must(build.Install(newAppScheme))
	utilruntime.Must(image.Install(newAppScheme))
	utilruntime.Must(network.Install(newAppScheme))
	utilruntime.Must(oauth.Install(newAppScheme))
	utilruntime.Must(project.Install(newAppScheme))
	utilruntime.Must(quota.Install(newAppScheme))
	utilruntime.Must(route.Install(newAppScheme))
	utilruntime.Must(security.Install(newAppScheme))
	utilruntime.Must(template.Install(newAppScheme))
	utilruntime.Must(user.Install(newAppScheme))
}
