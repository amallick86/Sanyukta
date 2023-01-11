package util

import (
	"database/sql"
	"time"
)

func CreateNullString(valid bool, value string) sql.NullString {
	nullString := sql.NullString{
		Valid:  valid,
		String: value,
	}
	return nullString
}

func CreateNullBool(valid bool, value bool) sql.NullBool {
	return sql.NullBool{
		Valid: valid,
		Bool:  value,
	}
}

func CreateNUllTime(valid bool, value time.Time) sql.NullTime {
	return sql.NullTime{
		Valid: valid,
		Time:  value,
	}
}
func CreateNullInt16(valid bool, value int16) sql.NullInt16 {
	return sql.NullInt16{
		Valid: valid,
		Int16: value,
	}
}
func CreateNullInt32(valid bool, value int32) sql.NullInt32 {
	return sql.NullInt32{
		Valid: valid,
		Int32: value,
	}
}
func CreateNullInt64(valid bool, value int64) sql.NullInt64 {
	return sql.NullInt64{
		Valid: valid,
		Int64: value,
	}
}
