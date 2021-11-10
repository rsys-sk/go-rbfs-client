/*
 * Copyright (C) 2021, RtBrick, Inc.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package commons

import (
	"net/http"
	"net/url"

	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/state"
)

// GetAPIClient creates a new API client for the given endpoint.
func GetAPIClient(client *http.Client, endpoint *url.URL) *state.APIClient {
	config := state.NewConfiguration()
	config.BasePath = endpoint.String()
	config.Host = endpoint.Host
	config.HTTPClient = client
	config.UserAgent = "Diagnostic Actor"
	return state.NewAPIClient(config)
}
