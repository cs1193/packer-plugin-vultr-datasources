package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
)

var (
	Version           = "0.0.0"
	VersionPrerelease = "dev"
)

var PluginVersion *version.PluginVersion

func init() {
	PluginVersion = version.InitializePluginVersion(Version, VersionPrerelease)
}

func main() {
	pps := plugin.NewSet()
	pps.SetVersion(PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
