package database

import (
	"reflect"
	"strings"
	"unicode"
)

func GetSelectFields[T any](item *T) []string {
	var fields []string
	v := reflect.ValueOf(item).Elem()
	t := reflect.TypeOf(item).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if strings.ToLower(fieldType.Name) == "password" {
			continue
		}

		if !isZeroValue(field) {
			dbName := getDBFieldName(fieldType)
			fields = append(fields, dbName)
		}
	}

	return fields
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface:
		return v.IsNil()
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}

func getDBFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("gorm")
	if tag != "" {
		parts := strings.Split(tag, ";")
		for _, part := range parts {
			if strings.HasPrefix(part, "column:") {
				return strings.TrimPrefix(part, "column:")
			}
		}
	}

	return toSnakeCase(field.Name)
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
