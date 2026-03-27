package types

import (
"encoding/json"
"testing"
)

func TestChargingSchedulePeriodLimitUnmarshal(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  float64
	}{
		{"numeric", `{"startPeriod":0,"limit":14.0}`, 14.0},
		{"string", `{"startPeriod":0,"limit":"14.0"}`, 14.0},
		{"string_int", `{"startPeriod":0,"limit":"6"}`, 6.0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
var p ChargingSchedulePeriod
if err := json.Unmarshal([]byte(tc.input), &p); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if p.Limit != tc.want {
				t.Errorf("Limit = %v, want %v", p.Limit, tc.want)
			}
		})
	}
}
