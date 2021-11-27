// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen_deps.go; DO NOT EDIT.

package integration

import (
	// And depend on a bunch of tailscaled innards, for Go's test caching.
	// Otherwise cmd/go never sees that we depend on these packages'
	// transitive deps when we run "go install tailscaled" in a child
	// process and can cache a prior success when a dependency changes.
	_ "inet.af/netaddr"
	_ "tailscale.com/chirp"
	_ "tailscale.com/derp/derphttp"
	_ "tailscale.com/ipn"
	_ "tailscale.com/ipn/ipnserver"
	_ "tailscale.com/logpolicy"
	_ "tailscale.com/logtail"
	_ "tailscale.com/net/dns"
	_ "tailscale.com/net/interfaces"
	_ "tailscale.com/net/netns"
	_ "tailscale.com/net/portmapper"
	_ "tailscale.com/net/proxymux"
	_ "tailscale.com/net/socks5/tssocks"
	_ "tailscale.com/net/tshttpproxy"
	_ "tailscale.com/net/tstun"
	_ "tailscale.com/paths"
	_ "tailscale.com/safesocket"
	_ "tailscale.com/tailcfg"
	_ "tailscale.com/types/flagtype"
	_ "tailscale.com/types/key"
	_ "tailscale.com/types/logger"
	_ "tailscale.com/util/clientmetric"
	_ "tailscale.com/util/multierr"
	_ "tailscale.com/util/osshare"
	_ "tailscale.com/version"
	_ "tailscale.com/version/distro"
	_ "tailscale.com/wgengine"
	_ "tailscale.com/wgengine/monitor"
	_ "tailscale.com/wgengine/netstack"
	_ "tailscale.com/wgengine/router"
)
