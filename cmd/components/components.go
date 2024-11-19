package components

import (
	"github.com/lburgazzoli/opendatahub-cli/cmd/components/list"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

const (
	CommandName  = "components"
	CommandAlias = "c"
)

func NewCmd(streams genericiooptions.IOStreams) *cobra.Command {
	cmd := cobra.Command{
		Use:     CommandName,
		Short:   CommandName,
		Aliases: []string{CommandAlias},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(list.NewCmd(streams))

	return &cmd
}
