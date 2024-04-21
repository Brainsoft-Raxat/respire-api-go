package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func CreateUpdateMap(model interface{}) (map[string]interface{}, error) {
	update := make(map[string]interface{})
	v := reflect.ValueOf(model)

	// Handle pointers by getting the value they point to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure the provided interface is a struct
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("createUpdateMap only accepts structs; got %s", v.Kind())
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("firestore")

		// Ignore fields with '-' or without 'firestore' tag
		if tag == "-" || tag == "" {
			continue
		}

		// Handle omitempty: skip zero values
		if strings.Contains(tag, "omitempty") {
			tag = strings.Split(tag, ",")[0] // Remove options
			if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
				continue
			}
		}

		update[tag] = field.Interface()
	}

	return update, nil
}
