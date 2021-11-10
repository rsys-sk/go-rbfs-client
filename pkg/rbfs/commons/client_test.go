/*
 * Copyright (C) 2021, RtBrick, Inc.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package commons

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAPIClient(t *testing.T) {
	tests := []struct {
		name     string
		client   *http.Client
		endpoint *url.URL
	}{
		{
			client:   http.DefaultClient,
			endpoint: mustParse(t, "http://192.168.0.1/path"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanics(t, func() { GetAPIClient(tt.client, tt.endpoint) })
		})
	}
}
