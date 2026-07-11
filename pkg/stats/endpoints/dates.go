package endpoints

import (
	"fmt"
	"strings"
	"time"
)

// parseBirthDate parses the raw BIRTHDATE/BIRTH_DATE string returned by the
// NBA Stats API into a time.Time. The upstream API is inconsistent about the
// format it uses across endpoints:
//   - CommonPlayerInfo/CommonPlayerInfoV2: "2006-01-02T15:04:05" (no timezone)
//   - CommonTeamRoster/CommonTeamRosterV2: "JAN 02, 2006" (uppercase month)
func parseBirthDate(raw string) (time.Time, error) {
	if raw == "" {
		return time.Time{}, fmt.Errorf("parse birth date: empty value")
	}

	if t, err := time.Parse("2006-01-02T15:04:05", raw); err == nil {
		return t, nil
	}

	if t, err := time.Parse("Jan 02, 2006", normalizeMonthCase(raw)); err == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("parse birth date %q: unrecognized format", raw)
}

// normalizeMonthCase converts the leading month token of a date string (e.g.
// "APR 17, 2001") to the mixed case Go's reference layout expects ("Apr 17, 2001").
func normalizeMonthCase(s string) string {
	month, rest, found := strings.Cut(s, " ")
	if !found || month == "" {
		return s
	}
	return strings.ToUpper(month[:1]) + strings.ToLower(month[1:]) + " " + rest
}

// DateOfBirth parses the player's birth date.
func (p PlayerInfo) DateOfBirth() (time.Time, error) {
	return parseBirthDate(p.Birthdate)
}

// DateOfBirth parses the player's birth date.
func (p CommonPlayerInfoV2CommonPlayerInfo) DateOfBirth() (time.Time, error) {
	return parseBirthDate(p.BIRTHDATE)
}

// DateOfBirth parses the player's birth date.
func (r CommonTeamRosterCommonTeamRoster) DateOfBirth() (time.Time, error) {
	return parseBirthDate(r.BIRTH_DATE)
}

// DateOfBirth parses the player's birth date.
func (r CommonTeamRosterV2CommonTeamRoster) DateOfBirth() (time.Time, error) {
	return parseBirthDate(r.BIRTH_DATE)
}
