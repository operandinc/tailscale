// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
	"tailscale.com/ipn/ipnstate"
	"tailscale.com/version"
)

var versionCmd = &ffcli.Command{
	Name:       "version",
	ShortUsage: "version [flags]",
	ShortHelp:  "Print Tailscale version",
	FlagSet: (func() *flag.FlagSet {
		fs := newFlagSet("version")
		fs.BoolVar(&versionArgs.daemon, "daemon", false, "also print local node's daemon version")
		fs.BoolVar(&versionArgs.json, "json", false, "output in JSON format")
		return fs
	})(),
	Exec: runVersion,
}

var versionArgs struct {
	daemon bool // also check local node's daemon version
	json   bool
}

func runVersion(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many non-flag arguments: %q", args)
	}
	var err error
	var st *ipnstate.Status

	if versionArgs.daemon {
		st, err = localClient.StatusWithoutPeers(ctx)
		if err != nil {
			return err
		}
	}

	if versionArgs.json {
		m := version.GetMeta()
		if st != nil {
			m.DaemonLong = st.Version
		}
		e := json.NewEncoder(os.Stdout)
		e.SetIndent("", "\t")
		return e.Encode(m)
	}

	if st == nil {
		outln(version.String())
		return nil
	}
	printf("Client: %s\n", version.String())
	printf("Daemon: %s\n", st.Version)
	return nil
}
