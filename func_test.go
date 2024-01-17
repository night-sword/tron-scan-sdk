package tronscan

import (
	"reflect"
	"testing"
)

func Test_toStringMap(t *testing.T) {
	tests := []struct {
		name    string
		fields  *ListTrc20TransfersWithStatusRequest
		wantM   map[string]string
		wantErr bool
	}{
		{
			"case.1",
			&ListTrc20TransfersWithStatusRequest{Start: 1, Limit: 10, Trc20Id: "Trc20Id", Address: "Address", Direction: 0, DbVersion: 0, Reverse: false},
			map[string]string{"start": "1", "limit": "10", "trc20Id": "Trc20Id", "address": "Address"},
			false,
		},
		{
			"case.2",
			&ListTrc20TransfersWithStatusRequest{Start: 1, Limit: 10, Trc20Id: "Trc20Id", Address: "Address", Direction: 0, DbVersion: 0, Reverse: true},
			map[string]string{"start": "1", "limit": "10", "trc20Id": "Trc20Id", "address": "Address", "reverse": "true"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM := toStringMap(tt.fields)
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("ToMap() gotM = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
