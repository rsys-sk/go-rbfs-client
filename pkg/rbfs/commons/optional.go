/*
 * Copyright (C) 2021, RtBrick, Inc.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package commons

import (
	"net"

	"github.com/antihax/optional"
)

func OptionalIP(ipAddress net.IP) optional.String {
	if ipAddress == nil {
		return optional.EmptyString()
	}
	return optional.NewString(ipAddress.String())
}

func OptionalString(s string) optional.String {
	if s == "" {
		return optional.EmptyString()
	}
	return optional.NewString(s)
}

func OptionalInt32(i int32) optional.Int32 {
	if i == 0 {
		return optional.EmptyInt32()
	}
	return optional.NewInt32(i)
}

func OptionalFloat32(f float32) optional.Float32 {
	if f == 0 {
		return optional.EmptyFloat32()
	}
	return optional.NewFloat32(f)
}
