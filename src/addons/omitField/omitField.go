package omitField

import (
	"reflect"
	"strings"
)

type (
	OmitFields struct{}
)

var (
	tagName = "groups"
)

func (r *OmitFields) ShowFieleds(obj interface{}, acTags []string) {
	sv := reflect.ValueOf(obj).Elem()
	st := sv.Type()
	if sv.Kind() == reflect.Struct {
		for i := 0; i < st.NumField(); i++ {
			fieldVal := sv.Field(i)
			if fieldVal.CanSet() {
				tagStr := st.Field(i).Tag.Get(tagName)
				if len(tagStr) == 0 {
					continue
				}
				tagList := strings.Split(strings.Replace(tagStr, " ", "", -1), ",")
				if !containCheck(tagList, acTags) {
					fieldVal.Set(reflect.Zero(fieldVal.Type()))
				}
			}
		}
	}
}

func containCheck(arr1 []string, arr2 []string) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}
	return false
}
