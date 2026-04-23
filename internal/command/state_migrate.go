// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform/internal/command/arguments"
	"github.com/hashicorp/terraform/internal/command/views"
)

// StateMigrateCommand is a Command implementation that migrates
// the state file from one location to another
type StateMigrateCommand struct {
	Meta
}

func (c *StateMigrateCommand) Run(rawArgs []string) int {
	// Parse and apply global view arguments
	common, rawArgs := arguments.ParseView(rawArgs)
	c.Meta.View.Configure(common)

	args, diags := arguments.ParseStateMigrate(rawArgs)

	stateMigrate := views.NewStateMigrate(args.ViewType, c.View)

	if diags.HasErrors() {
		stateMigrate.Diagnostics(diags)
		return 1
	}

	c.Meta.input = args.InputEnabled

	// stateMigrate.Printf("migrating from %s to %s", "source", "destination")

	// TODO: implement
	// parsedArgs.SourceLockFilePath
	// parsedArgs.DestinationLockFilePath
	// parsedArgs.Upgrade
	// parsedArgs.Input

	diags = diags.Append(errors.New("Not implemented yet"))
	stateMigrate.Diagnostics(diags)
	return 1
}

// terraform -input=false state migrate
func (c *StateMigrateCommand) Help() string {
	helpText := `
Usage: terraform [global options] state migrate [options]

  Migrate state from source declared in the migration configuration (*.tfmigrate.hcl)
  to the destination declared in the root module (*.tf).

  An error will be returned if the migration fails, e.g. if the state
  is inaccessible or the migration configuration is invalid.

Options:

  -source-provider-lock-file	   Path to a provider lock file for the source provider.

  -destination-provider-lock-file  Path to a provider lock file for the destination provider.

  -upgrade  					   Trigger upgrade of the provider.
  
  -input  					   	   TODO
`
	return strings.TrimSpace(helpText)
}

func (c *StateMigrateCommand) Synopsis() string {
	return "Migrate the state from one location to another"
}
