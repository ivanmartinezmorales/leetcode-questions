package utils

import (
	"errors"
	"math"
	"reflect"
)

// Contains returns true if the slice contains the value
func Contains(a []int, b int) bool {
	for _, k := range a {
		if k == b {
			return true
		}
	}
	return false
}

// MaxInt gets the margest int in the bunch
func MaxInt(a ...int) int {
	switch len(a) {
	case 0:
		return 0
	case 1:
		return a[0]
	case 2:
		if a[0] > a[1] {
			return a[0]
		}
		return a[1]
	default:
		m := math.MinInt32
		for _, k := range a {
			if k > m {
				m = k
			}
		}
		return m
	}
}

// MinInt returns the smallest int in the bunch.
func MinInt(a ...int) int {
	switch len(a) {
	case 0:
		return 0
	case 1:
		return a[0]
	case 2:
		if a[0] > a[1] {
			return a[1]
		}
		return a[0]
	default:
		m := math.MaxInt32
		for _, k := range a {
			if k < m {
				m = k
			}
		}
		return m
	}
}

// InsertAtSlice inserts an element to a slice in the specified index
func InsertAtSlice(slice interface{}, index int, value interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)

	if v.Kind() != reflect.Slice {
		return nil, errors.New("this is not a slice")
	}

	if index < 0 || index > v.Len() || reflect.TypeOf(slice).Elem() != reflect.ValueOf(value) {
		return nil, errors.New("param is invalid")
	}

	t := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	if index == v.Len() {
		t = reflect.AppendSlice(t, v.Slice(0, v.Len()))
		t = reflect.Append(t, reflect.ValueOf(value))
		return t.Interface(), nil
	}

	t = reflect.AppendSlice(t, v.Slice(0, v.Len()))
	t = reflect.AppendSlice(t, v.Slice(index, v.Len()))
	t.Index(index).Set(reflect.ValueOf(value))
	return t.Interface(), nil
}
