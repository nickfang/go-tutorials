package main

import "testing"

func TestGetDaysInMonth(t *testing.T) {
	numDays := GetDaysInMonth(2, 1800)
	if numDays != 28 {
		t.Error("should be 28")
	}
	numDays = GetDaysInMonth(2, 1900)
	if numDays != 28 {
		t.Error("should be 28")
	}
	numDays = GetDaysInMonth(2, 1901)
	if numDays != 28 {
		t.Error("should be 29")
	}
	numDays = GetDaysInMonth(2, 1902)
	if numDays != 28 {
		t.Error("should be 29")
	}
	numDays = GetDaysInMonth(2, 1903)
	if numDays != 28 {
		t.Error("should be 29")
	}
	numDays = GetDaysInMonth(2, 1904)
	if numDays != 29 {
		t.Error("should be 29")
	}
	numDays = GetDaysInMonth(2, 1908)
	if numDays != 29 {
		t.Error("should be 29")
	}
	numDays = GetDaysInMonth(2, 2000)
	if numDays != 29 {
		t.Error("should be 29")
	}
}
