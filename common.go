package utilities

import (
	"encoding/base64"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func TernaryOperator[A any](statement bool, valueIfTrue A, valueIfFalse A) A {
	if statement {
		return valueIfTrue
	}

	return valueIfFalse
}

func NewPointer[A any](a A) *A {
	return &a
}

func UnwrapPointer[A any](a *A) A {
	if a == nil {
		return *new(A)
	}

	return *a
}

func Map[A any, B any](f func(a A) B, amap []A) []B {
	bmap := make([]B, len(amap))
	for k, v := range amap {
		bmap[k] = f(v)
	}
	return bmap
}

func MapPointer[A any, B any](f func(a A) B, pointer *A) *B {
	if pointer == nil {
		return nil
	}

	result := f(*pointer)
	return &result
}

func Member[A comparable](a A, amap []A) bool {
	for _, v := range amap {
		if v == a {
			return true
		}
	}
	return false
}

func Deduplicate[A comparable](alist []A) []A {
	allKeys := make(map[A]bool)
	list := []A{}
	for _, item := range alist {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func GetDuplicateItems[A comparable](alist []A) []A {
	allKeys := make(map[A]bool)
	list := []A{}
	for _, item := range alist {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			continue
		}

		list = append(list, item)
	}
	return list
}

func Filter[A any](f func(a A) bool, amap []A) []A {
	res := make([]A, 0)
	for _, v := range amap {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func ToPgType[T any, P any](pointer *P, constructor func(*P) T) T {
	return constructor(pointer)
}

func NewPgText(value *string) pgtype.Text {
	if value != nil {
		return pgtype.Text{String: *value, Valid: true}
	}
	return pgtype.Text{String: "", Valid: false}
}

func NewPgInt8(value *int64) pgtype.Int8 {
	if value != nil {
		return pgtype.Int8{Int64: *value, Valid: true}
	}
	return pgtype.Int8{Valid: false}
}

func NewPgInt4(value *int32) pgtype.Int4 {
	if value != nil {
		return pgtype.Int4{Int32: *value, Valid: true}
	}
	return pgtype.Int4{Valid: false}
}

func NewPgFloat8(value *float64) pgtype.Float8 {
	if value != nil {
		return pgtype.Float8{Float64: *value, Valid: true}
	}
	return pgtype.Float8{Valid: false}
}

func NewPgBool(value *bool) pgtype.Bool {
	if value != nil {
		return pgtype.Bool{Bool: *value, Valid: true}
	}
	return pgtype.Bool{Valid: false}
}

func NewPgTimestamptz(value *time.Time) pgtype.Timestamptz {
	if value != nil {
		return pgtype.Timestamptz{Time: *value, Valid: true}
	}
	return pgtype.Timestamptz{Valid: false}
}

func FromPgType[T any, P any](value T, constructor func(T) *P) *P {
	return constructor(value)
}

func PgTypeToString(value pgtype.Text) *string {
	if value.Valid {
		return &value.String
	}
	return nil
}

func PgTypeToInt64(value pgtype.Int8) *int64 {
	if value.Valid {
		return &value.Int64
	}
	return nil
}

func PgTypeToInt32(value pgtype.Int4) *int32 {
	if value.Valid {
		return &value.Int32
	}
	return nil
}

func PgTypeToFloat64(value pgtype.Float8) *float64 {
	if value.Valid {
		return &value.Float64
	}
	return nil
}

func PgTypeToBool(value pgtype.Bool) *bool {
	if value.Valid {
		return &value.Bool
	}
	return nil
}

func PgTypeToTime(value pgtype.Timestamptz) *time.Time {
	if value.Valid {
		return &value.Time
	}
	return nil
}

func ToTimestamppb(val *time.Time) *timestamppb.Timestamp {
	if val == nil {
		return nil
	}
	return timestamppb.New(*val)
}

func FromTimestamppb(value *timestamppb.Timestamp) *time.Time {
	if value == nil {
		return nil
	}
	return NewPointer(value.AsTime())
}

func TimestamppbToTime(val *timestamppb.Timestamp) *time.Time {
	if val == nil {
		return nil
	}
	return NewPointer(val.AsTime())
}

func Base64Decode(base64Str string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
