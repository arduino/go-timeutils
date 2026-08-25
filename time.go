// This file is part of go-timeutils
//
// SPDX-FileCopyrightText: Arduino s.r.l. and/or its affiliated companies
// SPDX-License-Identifier: GPL-2.0-or-later

package timeutils

import "time"

// TimezoneOffsetNoDST returns the timezone offset without the DST component
func TimezoneOffsetNoDST(t time.Time) int {
	_, winterOffset := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location()).Zone()
	_, summerOffset := time.Date(t.Year(), 7, 1, 0, 0, 0, 0, t.Location()).Zone()
	if winterOffset > summerOffset {
		winterOffset, summerOffset = summerOffset, winterOffset
	}
	return winterOffset
}

// DaylightSavingsOffset returns the DST offset of the specified time
func DaylightSavingsOffset(t time.Time) int {
	_, offset := t.Zone()
	return offset - TimezoneOffsetNoDST(t)
}

// LocalUnix returns the unix timestamp of the specified time with the
// local timezone offset and DST added
func LocalUnix(t time.Time) int64 {
	_, offset := t.Zone()
	return t.Unix() + int64(offset)
}
