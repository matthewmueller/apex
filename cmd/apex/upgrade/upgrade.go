// Package upgrade installs the latest stable binary of Apex.
package upgrade

import (
	"github.com/tj/cobra"

	"github.com/matthewmueller/apex/cmd/apex/root"
	"github.com/matthewmueller/apex/cmd/apex/version"
	"github.com/matthewmueller/apex/upgrade"
)

// Command config.
var Command = &cobra.Command{
	Use:              "upgrade",
	Short:            "Upgrade apex to the latest stable release",
	PersistentPreRun: root.PreRunNoop,
	RunE:             run,
}

// Initialize.
func init() {
	root.Register(Command)
}

// Run command.
func run(c *cobra.Command, args []string) error {
	return upgrade.Upgrade(version.Version)
}
