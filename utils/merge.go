package utils

import "reflect"

func Merge(dst, src interface{}) {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dst)
	if s.Kind() != reflect.Ptr || d.Kind() != reflect.Ptr {
		return
	}
	for i := 0; i < s.Elem().NumField(); i++ {
		v := s.Elem().Field(i)
		fieldName := s.Elem().Type().Field(i).Name
		skip := s.Elem().Type().Field(i).Tag.Get("structs")
		if skip == "-" {
			continue
		}
		if v.Kind() > reflect.Float64 &&
			v.Kind() != reflect.String &&
			v.Kind() != reflect.Struct &&
			v.Kind() != reflect.Ptr &&
			v.Kind() != reflect.Slice {
			continue
		}
		if v.Kind() == reflect.Ptr {
			if !v.IsNil() {

				if d.Elem().FieldByName(fieldName).Kind() == reflect.Ptr {
					d.Elem().FieldByName(fieldName).Set(v)
					continue
				}
				f := d.Elem().FieldByName(fieldName)
				if f.IsValid() {
					f.Set(v.Elem())
				}
			}
			continue
		}
		d.Elem().FieldByName(fieldName).Set(v)
	}
}
