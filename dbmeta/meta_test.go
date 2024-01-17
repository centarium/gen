package dbmeta

import "testing"

func TestIsIntType(t *testing.T) {
	tests := []struct {
		fieldInfo *FieldInfo
		expected  bool
	}{
		{&FieldInfo{ProtobufType: "int32"}, true},
		{&FieldInfo{ProtobufType: "int64"}, true},
		{&FieldInfo{ProtobufType: "uint32"}, true},
		{&FieldInfo{ProtobufType: "uint64"}, true},
		{&FieldInfo{ProtobufType: "sint32"}, true},
		{&FieldInfo{ProtobufType: "sint64"}, true},
		{&FieldInfo{ProtobufType: "fixed32"}, true},
		{&FieldInfo{ProtobufType: "fixed64"}, true},
		{&FieldInfo{ProtobufType: "sfixed32"}, true},
		{&FieldInfo{ProtobufType: "sfixed64"}, true},
		{&FieldInfo{ProtobufType: "float32"}, false},
		{&FieldInfo{ProtobufType: "float64"}, false},
		{&FieldInfo{ProtobufType: "string"}, false},
		{&FieldInfo{ProtobufType: "bool"}, false},
		{&FieldInfo{ProtobufType: ""}, false},
	}

	for _, test := range tests {
		result := test.fieldInfo.IsIntType()
		if result != test.expected {
			t.Errorf("For type %s, expected %t, but got %t", test.fieldInfo.ProtobufType, test.expected, result)
		}
	}
}

type colMetaForTest struct {
	columnMeta
}

func (c *colMetaForTest) IsRequired() bool {
	return true
}

func TestGetFieldTags(t *testing.T) {

	tests := []struct {
		name           string
		fieldInfo      FieldInfo
		expectedResult string
	}{
		{
			name: "PrimaryInt",
			fieldInfo: FieldInfo{
				ProtobufType: "int64",
				ColumnMeta: &colMetaForTest{
					columnMeta{
						isPrimaryKey: true,
						columnType:   "Int8",
					},
				},
			},
			expectedResult: "[(validate.rules).int64.gt = 0,(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {type: INTEGER}]",
		},
		{
			name: "PrimaryUUID",
			fieldInfo: FieldInfo{
				ProtobufType: "string",
				ColumnMeta: &colMetaForTest{
					columnMeta{
						isPrimaryKey: true,
						columnType:   "uUiD",
					},
				},
			},
			expectedResult: "[(validate.rules).string.min_len = 1]",
		},
		{
			name: "PrimaryString",
			fieldInfo: FieldInfo{
				ProtobufType: "string",
				ColumnMeta: &colMetaForTest{
					columnMeta{
						isPrimaryKey: true,
						columnType:   "text",
					},
				},
			},
			expectedResult: "[(validate.rules).string.min_len = 1]",
		},
		{
			name: "StringNonZero",
			fieldInfo: FieldInfo{
				ProtobufType: "string",
				ColumnMeta: &colMetaForTest{
					columnMeta{Check: StringNonZero},
				},
			},
			expectedResult: "[(validate.rules).string.min_len = 1]",
		},
		{
			name: "NumberNonZero",
			fieldInfo: FieldInfo{
				ProtobufType: "int32",
				ColumnMeta: &colMetaForTest{
					columnMeta{Check: NumberNonZero},
				},
			},
			expectedResult: "[(validate.rules).int32.gt = 0,(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {type: INTEGER}]",
		},
		{
			name: "NoTags",
			fieldInfo: FieldInfo{
				ProtobufType: "float64",
				ColumnMeta:   &colMetaForTest{},
			},
			expectedResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fieldInfo.GetFieldTags()
			if result != tt.expectedResult {
				t.Errorf("Expected: %s, Got: %s", tt.expectedResult, result)
			}
		})
	}
}
