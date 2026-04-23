// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package arguments

import (
	"github.com/hashicorp/terraform/internal/tfdiags"
)

// StateMigrate represents the command-line arguments for the state migrate command.
type StateMigrate struct {
	SourceLockFilePath      string
	DestinationLockFilePath string
	Upgrade                 bool
	InputEnabled            bool

	ViewType ViewType
}

// ParseStateMigrate processes CLI arguments, returning a StateMigrate value and
// diagnostics. If errors are encountered, a StateMigrate value is still returned
// representing the best effort interpretation of the arguments.
func ParseStateMigrate(args []string) (*StateMigrate, tfdiags.Diagnostics) {
	var diags tfdiags.Diagnostics
	migrate := &StateMigrate{
		ViewType: ViewHuman,
	}

	// TODO
	defaultDstLockFilePath := ""

	var srcLockFilePath, dstLockFilePath string
	var upgrade, input bool
	cmdFlags := defaultFlagSet("state migrate")
	cmdFlags.StringVar(&srcLockFilePath, "source-provider-lock-file", "", "Path to a provider lock file for the source provider.")
	cmdFlags.StringVar(&dstLockFilePath, "destination-provider-lock-file", defaultDstLockFilePath, "Path to a provider lock file for the destination provider.")
	cmdFlags.BoolVar(&upgrade, "upgrade", false, "TODO")
	cmdFlags.BoolVar(&input, "input", true, "TODO")

	if err := cmdFlags.Parse(args); err != nil {
		diags = diags.Append(tfdiags.Sourceless(
			tfdiags.Error,
			"Failed to parse command-line flags",
			err.Error(),
		))
	}

	migrate.SourceLockFilePath = srcLockFilePath
	migrate.DestinationLockFilePath = dstLockFilePath
	migrate.Upgrade = upgrade
	migrate.InputEnabled = input

	return migrate, diags
}
