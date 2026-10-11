// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package libtailscale

import "tailscale.com/envknob"

// Set by the private build invocation, not stored credentials. Leaving them
// empty produces the normal native transport. Login must use wssControlURL.
var wssControlURL string
var wssEndpoint string

func configureWSS() {
	if wssControlURL == "" || wssEndpoint == "" {
		return
	}
	envknob.Setenv("TS_CONTROL_WEBSOCKET", "true")
	envknob.Setenv("TS_HEADSCALE_CONTROL_URL", wssControlURL)
	envknob.Setenv("TS_HEADSCALE_CONTROL_ENDPOINT", wssEndpoint)
	envknob.Setenv("TS_HEADSCALE_DOH", "false")
	envknob.Setenv("TS_HEADSCALE_USE_HOST_ROUTES", "false")
}
