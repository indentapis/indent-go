package main

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	provider "go.indent.com/indent-go/provider/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		Debug:        os.Getenv("INDENT_PROVIDER_DEBUG") == "true",
		ProviderAddr: "terraform.io/indent-tf-testing/indent",
		ProviderFunc: provider.Provider,
	})
}
