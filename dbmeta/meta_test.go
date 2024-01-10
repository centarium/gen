package dbmeta

import "testing"

func TestIsIntType(t *testing.T) {
	tests := []struct {
		fieldType string
		expected  bool
	}{
		{"int32", true},
		{"int64", true},
		{"uint32", true},
		{"uint64", true},
		{"sint32", true},
		{"sint64", true},
		{"fixed32", true},
		{"fixed64", true},
		{"sfixed32", true},
		{"sfixed64", true},
		{"float32", false},
		{"float64", false},
		{"string", false},
		{"bool", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsIntType(test.fieldType)
		if result != test.expected {
			t.Errorf("For type %s, expected %t, but got %t", test.fieldType, test.expected, result)
		}
	}
}
