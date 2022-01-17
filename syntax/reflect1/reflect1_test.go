package reflect1_test

import (
	"reflect"
	"testing"
)

type User struct {
	Id   int
	Name string
}

func Test1(t *testing.T) {
	var u User
	fill(&u)
}

func fill(i interface{}) {
	v := reflect.ValueOf(i).Elem()
	for i := 0; i < v.NumField(); i++ {
		fi := v.Type().Field(i)
		if f := v.FieldByName(fi.Name); f.CanSet() {
			switch v.Kind() {
			case reflect.Int:
				f.Set(reflect.ValueOf(1))
			case reflect.String:
				f.Set(reflect.ValueOf("vv"))
			}
		}
	}
}
