package datextractor

import (
	"testing"
	"time"
)

func TestDateExtractor_case1(t *testing.T) {

	rawDate, date, success := Extract("[2001-01-01 01:00:00,007] XYZ")
	if rawDate != "2001-01-01 01:00:00,007" {
		t.Errorf("Wrong raw date: %s", rawDate)
	}

	if date != time.Date(2001, time.January, 1, 1, 0, 0, 7000000, time.UTC) {
		t.Errorf("Wrong date: %s", date)
	}

	if !success {
		t.Errorf("Wrong success status: %t", success)
	}
}

func TestDateExtractor_case2(t *testing.T) {

	rawDate, date, success := Extract("[2001-01-01 01:00:00.007] XYZ")
	if rawDate != "2001-01-01 01:00:00.007" {
		t.Errorf("Wrong raw date: %s", rawDate)
	}

	if date != time.Date(2001, time.January, 1, 1, 0, 0, 7000000, time.UTC) {
		t.Errorf("Wrong date: %s", date)
	}

	if !success {
		t.Errorf("Wrong success status: %t", success)
	}
}

func TestDateExtractor_case3(t *testing.T) {

	rawDate, date, success := Extract("2001-01-01T01:00:00.007 XYZ")
	if rawDate != "2001-01-01T01:00:00.007" {
		t.Errorf("Wrong raw date: %s", rawDate)
	}

	if date != time.Date(2001, time.January, 1, 1, 0, 0, 7000000, time.UTC) {
		t.Errorf("Wrong date: %s", date)
	}

	if !success {
		t.Errorf("Wrong success status: %t", success)
	}
}

func TestDateExtractor_case4(t *testing.T) {

	rawDate, date, success := Extract("NO_DATE XYZ")
	if rawDate != "" {
		t.Errorf("Raw date should be empty")
	}

	if !date.IsZero() {
		t.Errorf("Date should be Zero")
	}

	if success {
		t.Errorf("Wrong success status: %t", success)
	}
}
