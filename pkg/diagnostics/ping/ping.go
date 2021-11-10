/*
 * Copyright (C) 2021, RtBrick, Inc.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package ping

import (
	"fmt"
	"net"
	"time"
)

type (
	// Ping contains all arguments to ping a destination IP address or hostname
	Ping struct {
		destinationIP   net.IP
		destinationA    string
		destinationAAAA string
		sourceInterface string
		sourceIP        net.IP
		instanceName    string
		count           int32
		interval        time.Duration
	}

	// Option applies a ping command argument
	Option func(*Ping) error
)

// NewPing creates a new ping command
func NewPing(options ...Option) (*Ping, error) {
	p := &Ping{
		instanceName: "default",
		count:        5,
		interval:     time.Second,
	}

	// Apply all given ping option
	for _, option := range options {
		if err := option(p); err != nil {
			return nil, err
		}
	}

	if p.destinationIP == nil && p.destinationA == "" && p.destinationAAAA == "" {
		return nil, fmt.Errorf("ping destination not specified")
	}

	return p, nil
}

// DestinationIP sets the ping destination IP address.
// Override destination host name settings, if any.
func DestinationIP(ipAddr net.IP) Option {
	return func(p *Ping) error {
		p.destinationIP = ipAddr
		p.destinationA = ""
		p.destinationAAAA = ""
		return nil
	}
}

// DestinationHostNameA sets the destination hostname that shall be translated to an IPv4 address (DNS A record)
func DestinationHostNameA(hostname string) Option {
	return func(p *Ping) error {
		p.destinationIP = nil
		p.destinationA = hostname
		p.destinationAAAA = ""
		return nil
	}
}

// DestinationHostNameAAAA sets the destination hostname that shall be translated to an IPv6 address (DNS AAAA record)
func DestinationHostNameAAAA(hostname string) Option {
	return func(p *Ping) error {
		p.destinationIP = nil
		p.destinationA = ""
		p.destinationAAAA = hostname
		return nil
	}
}

// SourceIP specifies the source IP address
func SourceIP(ipAddress net.IP) Option {
	return func(p *Ping) error {
		if ipAddress != nil {
			if p.sourceInterface != "" {
				return fmt.Errorf("source interface and source IP are mutual exclusive")
			}
			p.sourceIP = ipAddress
		}
		return nil
	}
}

// SourceInterface sets the ping source interface name.
// Source interface and source IP are mutual exclusive!
func SourceInterface(name string) Option {
	return func(p *Ping) error {
		if name != "" {
			if p.sourceIP != nil {
				return fmt.Errorf("source interface and source IP are mutual exclusive")
			}
			p.sourceInterface = name
		}
		return nil
	}
}

// Count sets the number of pings to be sent.
func Count(count int32) Option {
	return func(p *Ping) error {
		if count <= 0 {
			return fmt.Errorf("count value must be greater than 0")
		}

		const maxAllowedPings = 10
		if count > maxAllowedPings {
			return fmt.Errorf("count value must be less or equal than %d", maxAllowedPings)
		}

		p.count = count
		return nil
	}
}

// Interval sets the interval between two pings.
// The accepted interval range is between 1ms and 10 seconds.
func Interval(interval time.Duration) Option {
	return func(p *Ping) error {
		if interval < 1*time.Millisecond {
			return fmt.Errorf("interval must not be less than 1ms")
		}
		if interval > 10*time.Second {
			return fmt.Errorf("interval must not exceed 10s")
		}
		p.interval = interval
		return nil
	}
}

// InstanceName sets the routing instance name to run the ping command.
func InstanceName(instanceName string) Option {
	return func(p *Ping) error {
		if instanceName == "" {
			return fmt.Errorf("instance name must not be empty")
		}
		p.instanceName = instanceName
		return nil
	}
}

func (p *Ping) SourceInterface() string {
	return p.sourceInterface
}

func (p *Ping) SourceIP() net.IP {
	return p.sourceIP
}

func (p *Ping) Count() int32 {
	return p.count
}

func (p *Ping) Interval() time.Duration {
	return p.interval
}

func (p *Ping) DestinationIP() net.IP {
	return p.destinationIP
}

func (p *Ping) DestinationHostNameA() string {
	return p.destinationA
}

func (p *Ping) DestinationHostNameAAAA() string {
	return p.destinationAAAA
}
