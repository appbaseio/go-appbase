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
			if field.Type.Kind() == reflect.Slice {
				slice := reflect.ValueOf(obj).Elem().FieldByIndex([]int{i})
				if slice.Len() == 0 {
					emptyFields = append(emptyFields, field.Name)
				}
				for j := 0; j < slice.Len(); j++ {
					if slice.Index(j).Kind() == reflect.Ptr {
						err := validate(slice.Index(j).Interface())
						if err != nil {
							return err
						}
					} else if slice.Index(j).Interface() == reflect.Zero(slice.Index(j).Type()).Interface() {
						emptyFields = append(emptyFields, field.Name)
					}
				}
			} else if reflect.ValueOf(obj).Elem().FieldByIndex([]int{i}).Interface() == reflect.Zero(field.Type).Interface() {
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
