package shared

import "database/sql"

func HandleNullableIntColumn(data sql.NullInt32) *int {
	isNotNull := data.Valid
	var field int
	var fieldPtr *int
	if isNotNull {
		field = int(data.Int32)
		fieldPtr = &field
	} else {
		fieldPtr = nil
	}

	return fieldPtr
}

func HandleNullableInt64Column(data sql.NullInt64) *int64 {
	isNotNull := data.Valid
	var field int64
	var fieldPtr *int64
	if isNotNull {
		field = int64(data.Int64)
		fieldPtr = &field
	} else {
		fieldPtr = nil
	}

	return fieldPtr
}

func HandleNullableStringColumn(data sql.NullString) *string {
	isNotNull := data.Valid
	var field string
	var fieldPtr *string
	if isNotNull {
		field = data.String
		fieldPtr = &field
	} else {
		fieldPtr = nil
	}

	return fieldPtr
}
