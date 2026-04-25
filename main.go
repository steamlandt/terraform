// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Terraform is a tool for building, changing, and versioning infrastructure
// safely and efficiently. Configuration files describe to Terraform the
// components needed to run a single application or your entire datacenter.
// Terraform generates an execution plan describing what it will do to reach
// the desired state, and then executes it to build the described
// infrastructure. As the configuration changes, Terraform is able to
// determine what changed and create incremental execution plans which can
// be applied.
//
// For more information, see the README in this directory.
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/opentofu/opentofu/internal/logging"
)

func main() {
	// Set up logging first so that any early failures are properly logged.
	logging.SetOutput()

	// The actual command dispatch happens in the commands package. We just
	// set up the environment here.
	os.Exit(realMain())
}

func realMain() int {
	defer logging.PanicHandler()

	// On non-Windows systems, we need to make sure that we're not running
	// as root. This is a security measure to prevent accidental damage.
	if runtime.GOOS != "windows" {
		if os.Getuid() == 0 {
			fmt.Fprintf(os.Stderr, "Running Terraform as root is not recommended and may cause unexpected behavior.\n")
		}
	}

	// Run the CLI
	cli, err := NewCLI()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing CLI: %s\n", err)
		return 1
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err)
		return 1
	}

	return exitCode
}
