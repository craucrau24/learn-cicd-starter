package auth

import (
	"testing"
)

func CheckSuccess(success bool, err error) bool {
	return (success && (err == nil)) || (!success && (err != nil))
}

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		hname string
		hvalue string
		want struct {
			value string
			success bool
		}
	}{"missing header": {hname: "Foo", hvalue: "Bar", want: struct{value string; success bool}{value: "", success: false}},
	"wrong header (1field)": {hname: "Authorization", hvalue: "ApiKey", want: struct{value string; success bool}{value: "", success: false}},
	"wrong header (wrong auth type)": {hname: "Authorization", hvalue: "Foo Bar", want: struct{value string; success bool}{value: "", success: false}},
	"correct header": {hname: "Authorization", hvalue: "ApiKey FooBar", want: struct{value string; success bool}{value: "FooBar", success: true}},
}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			headers := map[string][]string {tc.hname: {tc.hvalue}}
			got, err := GetAPIKey(headers)
			if !CheckSuccess(tc.want.success, err) {
				t.Fatalf("Test %v: error - %v, success expected: %v", name, err, tc.want.success)
			}
			if got != tc.want.value {
				t.Fatalf("Test %v: got - %v, expected: %v", name, got, tc.want.value)
			}
		})
	}
}