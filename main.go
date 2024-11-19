package main

import (
	"github.com/lburgazzoli/opendatahub-cli/cmd"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"os"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-odh", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdOpenDataHub(genericiooptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	})
	
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
