package actions

import (
	"errors"
	"reflect"
	"strings"
)

func validate(obj interface{}) error {
	emptyFields := make([]string, 0)
	var field reflect.StructField
	for i := 0; i < reflect.TypeOf(obj).Elem().NumField(); i++ {
		field = reflect.TypeOf(obj).Elem().FieldByIndex([]int{i})
		if field.Tag.Get("validate") == "required" {
			if reflect.ValueOf(obj).Elem().FieldByIndex([]int{i}).Interface() == reflect.Zero(field.Type).Interface() {
				emptyFields = append(emptyFields, field.Name)
			}
		}
	}

	if len(emptyFields) > 0 {
		missingFields := strings.Join(emptyFields, ", ")
		return errors.New(strings.Join([]string{"The following fields are required:", missingFields}, " "))
	} else {
		return nil
	}
}
