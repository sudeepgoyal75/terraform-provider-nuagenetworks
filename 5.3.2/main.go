package main

import (
	"github.com/hashicorp/terraform/plugin"
	nuagenetworks "github.com/tpretz/terraform-provider-nuagenetworks/5.3.2/nuagenetworks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nuagenetworks.Provider})
}
