package main

import "github.com/eure/kamimai"

var (
	syncCmd = &Cmd{
		Name:  "sync",
		Usage: "apply all migrations",
		Run:   doSyncCmd,
	}
)

func doSyncCmd(cmd *Cmd, args ...string) error {

	// Sync all migrations
	return kamimai.Sync(config)
}
