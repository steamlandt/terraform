// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"testing"

	"github.com/hashicorp/cli"
)

func TestStateMigrate(t *testing.T) {
	ui := cli.NewMockUi()
	view, done := testView(t)
	c := &StateMigrateCommand{
		Meta: Meta{
			Ui:   ui,
			View: view,
		},
	}

	args := []string{}
	if code := c.Run(args); code != 0 {
		out := done(t)
		t.Fatalf("bad: %d\n\n%s", code, out.All())
	}
}
