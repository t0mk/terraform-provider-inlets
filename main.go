package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/t0mk/terraform-provider-inlets/inlets"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: inlets.Provider})
}
