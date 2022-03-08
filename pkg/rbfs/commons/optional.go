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

func OptionalInt(i int) optional.Int {
	if i == 0 {
		return optional.EmptyInt()
	}
	return optional.NewInt(i)
}

func OptionalInt32(i int32) optional.Int32 {
	if i == 0 {
		return optional.EmptyInt32()
	}
	return optional.NewInt32(i)
}

func OptionalInt64(i int64) optional.Int64 {
	if i == 0 {
		return optional.EmptyInt64()
	}
	return optional.NewInt64(i)
}

func OptionalFloat32(f float32) optional.Float32 {
	if f == 0 {
		return optional.EmptyFloat32()
	}
	return optional.NewFloat32(f)
}

func OptionalFloat64(f float64) optional.Float64 {
	if f == 0 {
		return optional.EmptyFloat64()
	}
	return optional.NewFloat64(f)
}
