package types

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type BaseCmdOptions struct {
	genericiooptions.IOStreams

	kubernetes kubernetes.Interface
	dynamic    dynamic.Interface
	discovery  discovery.DiscoveryInterface

	Flags      *genericclioptions.ConfigFlags
	ValidateFn func(_ *cobra.Command, _ []string) error
	RunFn      func(_ *cobra.Command, _ []string) error
}

func (o *BaseCmdOptions) complete(_ *cobra.Command, _ []string) error {
	cfg, err := o.Flags.ToRawKubeConfigLoader().ClientConfig()
	if err != nil {
		return err
	}

	o.kubernetes, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("unable to construct a Kubernetes client: %w", err)
	}

	o.dynamic, err = dynamic.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("unable to construct a Dynamic client: %w", err)
	}

	o.discovery, err = o.Flags.ToDiscoveryClient()
	if err != nil {
		return fmt.Errorf("unable to construct a Dynamic client: %w", err)
	}

	return nil
}

func (o *BaseCmdOptions) Run(cmd *cobra.Command, args []string) error {
	if err := o.complete(cmd, args); err != nil {
		return err
	}

	if o.ValidateFn == nil {
		if err := o.ValidateFn(cmd, args); err != nil {
			return err
		}
	}

	return o.RunFn(cmd, args)
}

func (o *BaseCmdOptions) Client() kubernetes.Interface {
	return o.kubernetes
}

func (o *BaseCmdOptions) Dynamic() dynamic.Interface {
	return o.dynamic
}

func (o *BaseCmdOptions) Discovery() discovery.DiscoveryInterface {
	return o.discovery
}
