package utils

import (
	"reflect"
	"testing"
)

func TestGetFactors(t *testing.T) {
	factors := GetFactors(4)
	expected := []int{1, 2, 4}
	if !reflect.DeepEqual(factors, expected) {
		t.Errorf("GetFactors(4) = %v; want %v", factors, expected)
	}
	factors = GetFactors(17)
	expected = []int{1, 17}
	if !reflect.DeepEqual(factors, expected) {
		t.Errorf("GetFactors(4) = %v; want %v", factors, expected)
	}
	factors = GetFactors(28)
	expected = []int{1, 2, 4, 7, 14, 28}
	if !reflect.DeepEqual(factors, expected) {
		t.Errorf("GetFactors(28) = %v; want %v", factors, expected)
	}
	factors = GetFactors(196)
	expected = []int{1, 2, 4, 7, 14, 28, 49, 98, 196}
	if !reflect.DeepEqual(factors, expected) {
		t.Errorf("GetFactors(196) = %v; want %v", factors, expected)
	}
}

func TestGetProperDivisors(t *testing.T) {
	divisors := GetProperDivisors(4)
	expected := []int{1, 2}
	if !reflect.DeepEqual(divisors, expected) {
		t.Errorf("GetProperDivisors(4) = %v; want %v", divisors, expected)
	}
	divisors = GetProperDivisors(17)
	expected = []int{1}
	if !reflect.DeepEqual(divisors, expected) {
		t.Errorf("GetProperDivisors(28) = %v; want %v", divisors, expected)
	}
	divisors = GetProperDivisors(28)
	expected = []int{1, 2, 4, 7, 14}
	if !reflect.DeepEqual(divisors, expected) {
		t.Errorf("GetProperDivisors(28) = %v; want %v", divisors, expected)
	}
	divisors = GetProperDivisors(196)
	expected = []int{1, 2, 4, 7, 14, 28, 49, 98}
	if !reflect.DeepEqual(divisors, expected) {
		t.Errorf("GetProperDivisors(196) = %v; want %v", divisors, expected)
	}
}
