// This file is part of go-timeutils
//
// SPDX-FileCopyrightText: Arduino s.r.l. and/or its affiliated companies
// SPDX-License-Identifier: GPL-2.0-or-later

package timeutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	cet, err := time.LoadLocation("CET")
	require.NoError(t, err)
	ast, err := time.LoadLocation("Australia/Sydney")
	require.NoError(t, err)

	firstJanuary2015CET := time.Date(2015, 1, 1, 0, 0, 0, 0, cet)
	require.Equal(t, int64(1420066800), firstJanuary2015CET.Unix())
	require.Equal(t, int64(1420066800+3600), LocalUnix(firstJanuary2015CET))
	require.Equal(t, 3600, TimezoneOffsetNoDST(firstJanuary2015CET))
	require.Equal(t, 0, DaylightSavingsOffset(firstJanuary2015CET))

	fall2015CET := time.Date(2015, 9, 23, 0, 0, 0, 0, cet)
	require.Equal(t, int64(1442959200), fall2015CET.Unix())
	require.Equal(t, int64(1442959200+3600+3600), LocalUnix(fall2015CET))
	require.Equal(t, 3600, TimezoneOffsetNoDST(fall2015CET))
	require.Equal(t, 3600, DaylightSavingsOffset(fall2015CET))

	firstJan2015AST := time.Date(2015, 1, 1, 0, 0, 0, 0, ast)
	require.Equal(t, int64(1420030800), firstJan2015AST.Unix())
	require.Equal(t, int64(1420030800+36000+3600), LocalUnix(firstJan2015AST))
	require.Equal(t, 36000, TimezoneOffsetNoDST(firstJan2015AST))
	require.Equal(t, 3600, DaylightSavingsOffset(firstJan2015AST))

	fall2015AST := time.Date(2015, 9, 23, 0, 0, 0, 0, ast)
	require.Equal(t, int64(1442930400), fall2015AST.Unix())
	require.Equal(t, int64(1442930400+36000), LocalUnix(fall2015AST))
	require.Equal(t, 36000, TimezoneOffsetNoDST(fall2015AST))
	require.Equal(t, 0, DaylightSavingsOffset(fall2015AST))
}
