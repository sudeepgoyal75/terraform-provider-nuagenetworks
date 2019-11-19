package main

import (
	"github.com/hashicorp/terraform/plugin"
	nuagenetworks "github.com/tpretz/terraform-provider-nuagenetworks/3.2.7/nuagenetworks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nuagenetworks.Provider})
}
