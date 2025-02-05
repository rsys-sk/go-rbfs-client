/*
 * Copyright (C) 2021, RtBrick, Inc.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package ping

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewPing(t *testing.T) {
	tests := []struct {
		name    string
		options []Option
		want    *Ping
		wantErr string
	}{
		{
			name: "defaults",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
			},
			want: &Ping{
				destinationA: "www.rtbrick.com",
				instanceName: "default",
				count:        5,
				interval:     time.Second,
				size:         56,
				ttl:          64,
			},
		}, {
			name: "options",
			options: []Option{
				DestinationIP(net.ParseIP("8.8.8.8")),
				SourceIP(net.ParseIP("192.0.2.1")), Count(8),
				Interval(5 * time.Second), InstanceName("instance"),
			},
			want: &Ping{
				destinationIP:   net.ParseIP("8.8.8.8"),
				sourceInterface: "",
				sourceIP:        net.ParseIP("192.0.2.1"),
				instanceName:    "instance",
				count:           8,
				interval:        5 * time.Second,
				size:            56,
				ttl:             64,
			},
		}, {
			name: "source interface",
			options: []Option{
				DestinationHostNameAAAA("www.rtbrick.com"),
				SourceInterface("ma1"),
			},
			want: &Ping{
				destinationAAAA: "www.rtbrick.com",
				sourceInterface: "ma1",
				instanceName:    "default",
				count:           5,
				interval:        time.Second,
				size:            56,
				ttl:             64,
			},
		}, {
			name: "source interface and source IP are mutual exclusive 1",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
				SourceInterface("ma1"),
				SourceIP(net.ParseIP("192.0.2.1")),
			},
			wantErr: "source interface and source IP are mutual exclusive",
		}, {
			name: "source interface and source IP are mutual exclusive 2",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
				SourceIP(net.ParseIP("192.0.2.1")),
				SourceInterface("ma1"),
			},
			wantErr: "source interface and source IP are mutual exclusive",
		}, {
			name: "count value must be greater than 0",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
				Count(0),
			},
			wantErr: "count value must be greater than 0",
		}, {
			name: "count value must be less or equal than 10k",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
				Count(10001),
			},
			wantErr: "count value must not be greater than 10000",
		}, {
			name: "interval must not be less than 1ms",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
				Interval(time.Nanosecond),
			},
			wantErr: "interval must not be less than 1ms",
		}, {
			name: "interval must not exceed 5s",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
				Interval(6 * time.Second),
			},
			wantErr: "interval must not exceed 5s",
		}, {
			name: "instance name must not be empty",
			options: []Option{
				DestinationHostNameA("www.rtbrick.com"),
				InstanceName(""),
			},
			wantErr: "instance name must not be empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPing(tt.options...)
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
