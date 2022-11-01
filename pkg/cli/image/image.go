package image

import (
	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
	ktemplates "k8s.io/kubectl/pkg/util/templates"

	"github.com/uccps-samples/oc/pkg/cli/image/append"
	"github.com/uccps-samples/oc/pkg/cli/image/extract"
	"github.com/uccps-samples/oc/pkg/cli/image/info"
	"github.com/uccps-samples/oc/pkg/cli/image/mirror"
	"github.com/uccps-samples/oc/pkg/cli/image/serve"
	cmdutil "github.com/uccps-samples/oc/pkg/helpers/cmd"
)

var (
	imageLong = ktemplates.LongDesc(`
		Manage images on Uccp

		These commands help you manage images on Uccp.`)
)

// NewCmdImage exposes commands for modifying images.
func NewCmdImage(f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	image := &cobra.Command{
		Use:   "image COMMAND",
		Short: "Useful commands for managing images",
		Long:  imageLong,
		Run:   kcmdutil.DefaultSubCommandRun(streams.ErrOut),
	}

	groups := ktemplates.CommandGroups{
		{
			Message: "View or copy images:",
			Commands: []*cobra.Command{
				info.NewInfo(streams),
				mirror.NewCmdMirrorImage(streams),
			},
		},
		{
			Message: "Advanced commands:",
			Commands: []*cobra.Command{
				serve.NewServe(streams),
				append.NewCmdAppendImage(streams),
				extract.NewExtract(streams),
			},
		},
	}
	groups.Add(image)
	cmdutil.ActsAsRootCommand(image, []string{"options"}, groups...)
	return image
}
