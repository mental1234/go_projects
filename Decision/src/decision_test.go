package main

import (
	"reflect"
	"testing"
)

func TestCondition(t *testing.T) {
	val1, val2 := reflect.TypeOf(5), reflect.TypeOf(8)

	if val1 != val2 {
		t.Errorf("Expected 2 integers. Got 0 or 1")
	}
}
