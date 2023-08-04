package vo

import (
	"fmt"
	"reflect"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

type CustomAttributes map[string]any

func NewCustomAttributes(customAttributes map[string]any) (CustomAttributes, domain.DomainError) {
	return CustomAttributes(customAttributes), nil
}

func (vo CustomAttributes) GetAttribute(attributeName string) (attributeValue any, find bool) {
	attributeValue, find = vo[attributeName]
	return attributeValue, find
}

func (vo CustomAttributes) Flatten() map[string]any {
	flattened := make(map[string]any)
	flatten("", vo, flattened)
	return flattened
}

func flatten(prefix string, m map[string]any, flattened map[string]any) {
	for k, v := range m {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Map:
			flatten(key, v.(map[string]any), flattened)
		case reflect.Slice:
			flattenArray(key, interfaceToSlice(v), flattened)
		default:
			flattened[key] = v
		}
	}
}

func flattenArray(prefix string, arr []any, flattened map[string]any) {
	for i, v := range arr {
		key := fmt.Sprintf("%s[%d]", prefix, i)

		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Map:
			flatten(key, v.(map[string]any), flattened)
		case reflect.Slice:
			flattenArray(key, interfaceToSlice(v), flattened)
		default:
			flattened[key] = v
		}
	}
}

func interfaceToSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
