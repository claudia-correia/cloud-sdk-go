// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package api

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const transportCastErrFmt = "transport: failed converting %T to *http.Transport"

// TransportConfig is meant to be used so an http.RoundTripper is constructed
// with the appropriate settings.
type TransportConfig struct {
	// When SkipTLSVerify the TLS verification is completely skipped.
	SkipTLSVerify bool

	// ErrorDevice where any error or notices will be sent.
	ErrorDevice io.Writer

	// Can enable a debug RoundTripper which dumps the request and responses to
	// the configured device.
	VerboseSettings
}

func newDefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

// NewTransport constructs a new http.RoundTripper from its config.
func NewTransport(rt http.RoundTripper, cfg TransportConfig) http.RoundTripper {
	if rt == nil {
		// Change this to use the new .Clone() method once Go 1.13+ is
		// released. See https://github.com/golang/go/issues/26013.
		rt = newDefaultTransport()
	}

	switch t := rt.(type) {
	case *http.Transport:
		if t.TLSClientConfig == nil {
			t.TLSClientConfig = new(tls.Config)
		}
		t.TLSClientConfig.InsecureSkipVerify = cfg.SkipTLSVerify
		rt = t
	case *DebugTransport:
		return t
	default:
		if cfg.ErrorDevice != nil {
			fmt.Fprintf(cfg.ErrorDevice, transportCastErrFmt, rt)
		}
	}

	if cfg.Verbose {
		return NewDebugTransport(rt, cfg.Device)
	}

	return rt
}
