package generics_and_reflection

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	declTagValueOmitempty = "omitempty"
	declTagNameProperties = "properties"
	serializeSeparator    = "="
)

func Serialize[T any](s T) string {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return ""
	}

	var serialized strings.Builder
	v := reflect.ValueOf(s)

	for fIndex := range t.NumField() {
		fieldType := t.Field(fIndex)
		fieldValue := v.Field(fIndex)

		var (
			fieldName string
			omitempty bool
		)

		if tag, ok := fieldType.Tag.Lookup(declTagNameProperties); ok {
			parts := strings.SplitN(tag, ",", 2)
			fieldName = parts[0]
			omitempty = len(parts) > 1 && parts[1] == declTagValueOmitempty
		} else {
			fieldName = strings.ToLower(fieldType.Name)
		}

		if !omitempty || !fieldValue.IsZero() {
			serialized.WriteString(fieldName + serializeSeparator + fmt.Sprintf("%v", fieldValue))
			serialized.WriteByte('\n')
		}
	}

	return strings.TrimSuffix(serialized.String(), "\n")
}
