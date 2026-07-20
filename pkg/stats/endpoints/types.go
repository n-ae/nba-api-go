package endpoints

import (
	"fmt"
	"reflect"
)

type resultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}

type rawStatsResponse struct {
	ResultSets []resultSet `json:"resultSets"`
}

// findResultSet returns the result set named name, keyed by its actual
// "name" field rather than its position in resultSets. NBA.com does not
// guarantee result-set order is stable; positional indexing
// (rawResp.ResultSets[0], [1], ...) silently reads the wrong result set
// into the wrong struct field if the order ever changes upstream.
func findResultSet(resultSets []resultSet, name string) (resultSet, bool) {
	for _, rs := range resultSets {
		if rs.Name == name {
			return rs, true
		}
	}
	return resultSet{}, false
}

// validateHeaders errors if a result set's actual column headers (as
// reported by NBA.com in this response) don't match the field order this
// generated code assumes when it indexes row[i] positionally. Without
// this, an NBA.com column reorder or insertion would silently shift every
// field after the change into the wrong struct field instead of failing
// loudly - "shifting columns" is exactly the failure mode this exists to
// turn into a caught error.
func validateHeaders(got, want []string) error {
	if len(got) != len(want) {
		return fmt.Errorf("expected %d columns, got %d: %v", len(want), len(got), got)
	}
	for i := range want {
		if got[i] != want[i] {
			return fmt.Errorf("column %d: expected %q, got %q: %v", i, want[i], got[i], got)
		}
	}
	return nil
}

// jsonTags returns v's exported struct fields' `json:"..."` tag values, in
// field declaration order - the order generated code assumes when it
// indexes a result set's row positionally. Deriving this by reflecting
// over the struct itself, rather than generating a second, separately
// maintained list of the same field names as string literals, keeps a
// single source of truth: the struct's own field order and json tags.
func jsonTags(v interface{}) []string {
	t := reflect.TypeOf(v)
	tags := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		tags = append(tags, t.Field(i).Tag.Get("json"))
	}
	return tags
}

// Type conversion helpers for parsing NBA API responses
func toInt(v interface{}) int {
	switch val := v.(type) {
	case float64:
		return int(val)
	case int:
		return val
	case string:
		return 0
	default:
		return 0
	}
}

func toFloat(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case int:
		return float64(val)
	case string:
		return 0.0
	default:
		return 0.0
	}
}

func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%.0f", val)
	case int:
		return fmt.Sprintf("%d", val)
	default:
		return ""
	}
}
