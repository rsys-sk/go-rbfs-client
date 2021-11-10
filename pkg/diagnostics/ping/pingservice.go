/*
 * Copyright (C) 2021, RtBrick, Inc.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package ping

import (
	"context"
	"math"
	"net/http"
	"net/url"

	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/commons"
	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/state"
)

var (
	// This function variable we need to be able to mock the API Service
	getActionsAPIFunc = getActionsAPI
)

type (
	defaultService struct {
		client *http.Client
	}

	ActionsAPI interface {
		Ping(ctx context.Context, localVarOptionals *state.ActionsApiPingOpts) (state.PingStatus, *http.Response, error)
	}

	// Service pings given destinations.
	Service interface {
		// Run executes the given ping.
		Run(commons.RbfsContext, *Ping) (state.PingStatus, error)

		// RunAll runs all given pings in parallel go routines.
		RunAll(commons.RbfsContext, ...*Ping) ([]state.PingStatus, error)
	}
)

func (s *defaultService) Run(ctx commons.RbfsContext, ping *Ping) (state.PingStatus, error) {
	api, err := s.getActionsAPI(ctx)
	if err != nil {
		return state.PingStatus{}, err
	}

	const scaleToMilliPrecision = 1000
	interval := float32(math.Round(ping.interval.Seconds()*scaleToMilliPrecision) / scaleToMilliPrecision)

	optionalPingPostArgs := &state.ActionsApiPingOpts{
		DestinationIp:   commons.OptionalIP(ping.destinationIP),
		DestinationAaaa: commons.OptionalString(ping.destinationAAAA),
		DestinationA:    commons.OptionalString(ping.destinationA),
		SourceIp:        commons.OptionalIP(ping.sourceIP),
		SourceIfl:       commons.OptionalString(ping.sourceInterface),
		Count:           commons.OptionalInt32(ping.count),
		Interval:        commons.OptionalFloat32(interval),
		InstanceName:    commons.OptionalString(ping.instanceName),
	}

	//nolint:bodyclose //generated code
	pingStatus, _, err := api.Ping(ctx, optionalPingPostArgs)
	if err != nil {
		return state.PingStatus{}, err
	}
	return pingStatus, nil
}

func (s *defaultService) RunAll(ctx commons.RbfsContext, pings ...*Ping) ([]state.PingStatus, error) {
	var r []state.PingStatus
	dataDataChannel := make(chan state.PingStatus)
	errChannel := make(chan error)

	c, cancle := context.WithCancel(ctx)
	defer cancle()

	rc := commons.MustRbfsContext(c)
	for _, ping := range pings {
		go func(p *Ping) {
			s, err := s.Run(rc, p)
			if err != nil {
				errChannel <- err
				return
			}
			dataDataChannel <- s
		}(ping)
	}

	var err error
	for range pings {
		select {
		case data := <-dataDataChannel:
			r = append(r, data)
		case e := <-errChannel:
			if err == nil {
				err = e
				cancle()
			}
		}
	}
	return r, err
}

func (s *defaultService) getActionsAPI(ctx commons.RbfsContext) (ActionsAPI, error) {
	endpoint, err := ctx.GetServiceEndpoint("opsd")
	if err != nil {
		return nil, err
	}

	return getActionsAPIFunc(s.client, endpoint)
}

func getActionsAPI(c *http.Client, endpoint *url.URL) (ActionsAPI, error) {
	client := commons.GetAPIClient(c, endpoint)
	return client.ActionsApi, nil
}

// NewPingService creates a new ping defaultService.
func NewPingService(client *http.Client) Service {
	return &defaultService{client}
}
