package cmd

import (
	"github.com/lburgazzoli/opendatahub-cli/cmd/components"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

const (
	CommandName       = "odh"
	CommandAnnotation = "kubectl odh"
)

func NewCmdOpenDataHub(streams genericiooptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:          CommandName,
		Short:        CommandName,
		SilenceUsage: true,
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: CommandAnnotation,
		},
		RunE: func(c *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(components.NewCmd(streams))

	return cmd
}
