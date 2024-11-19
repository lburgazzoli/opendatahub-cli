package list

import (
	"fmt"
	"github.com/lburgazzoli/opendatahub-cli/cmd/types"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/cli-runtime/pkg/printers"
	"strings"
)

const (
	CommandName  = "list"
	CommandAlias = "ls"
)

func NewCmd(streams genericiooptions.IOStreams) *cobra.Command {
	o := types.BaseCmdOptions{
		Flags:     genericclioptions.NewConfigFlags(true),
		IOStreams: streams,
	}

	o.RunFn = func(c *cobra.Command, args []string) error {
		w := printers.GetNewTabWriter(o.Out)

		lists, err := o.Discovery().ServerPreferredResources()
		if err != nil {
			return err
		}

		for _, list := range lists {
			if len(list.APIResources) == 0 {
				continue
			}

			gv, err := schema.ParseGroupVersion(list.GroupVersion)
			if err != nil {
				continue
			}

			for _, resource := range list.APIResources {
				if len(resource.Verbs) == 0 {
					continue
				}
				if gv.Group != "components.opendatahub.com" {
					continue
				}
			}
		}

		columnNames := []string{"NAME", "READY", "REASON"}
		if _, err := fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t")); err != nil {
			return err
		}

		return nil
	}

	cmd := cobra.Command{
		Use:     CommandName,
		Short:   CommandName,
		Aliases: []string{CommandAlias},
		RunE:    o.Run,
	}

	return &cmd
}
