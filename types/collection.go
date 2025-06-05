package types

import "reflect"

// Contains returns whether the collection contains the element.
// The collection supports array, slice, and map.
func Contains(collection interface{}, that interface{}) bool {
	targetValue := reflect.ValueOf(collection)
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == that {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(that)).IsValid() {
			return true
		}
	default:
		return false
	}

	return false
}
