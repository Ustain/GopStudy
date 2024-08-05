package models

type TestCase struct {
    FieldName   string `json:"field_name" validate:"required"`
    FieldLength int    `json:"field_length" validate:"gte=0"`
    FieldType   string `json:"field_type" validate:"required,oneof=int string bool"`
}

type APIRequest struct {
    TestCases []TestCase `json:"test_cases" validate:"required,dive"`
}
