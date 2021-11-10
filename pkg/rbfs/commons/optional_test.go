/*
 * Copyright (C) 2021, RtBrick, Inc.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package commons

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionalString(t *testing.T) {
	o := OptionalString("test")
	require.True(t, o.IsSet())
	require.Equal(t, "test", o.Value())

	o = OptionalString("")
	require.False(t, o.IsSet())
	require.Equal(t, "", o.Value())
}

func TestOptionalInt32(t *testing.T) {
	o := OptionalInt32(10)
	require.True(t, o.IsSet())
	require.Equal(t, int32(10), o.Value())

	o = OptionalInt32(0)
	require.False(t, o.IsSet())
	require.Equal(t, int32(0), o.Value())
}

func TestOptionalFloat32(t *testing.T) {
	o := OptionalFloat32(10)
	require.True(t, o.IsSet())
	require.Equal(t, float32(10), o.Value())

	o = OptionalFloat32(0)
	require.False(t, o.IsSet())
	require.Equal(t, float32(0), o.Value())
}
