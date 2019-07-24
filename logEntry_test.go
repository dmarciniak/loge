package loge

import (
	"testing"
	"time"
)

func TestEmptyLogEntry(t *testing.T) {
	if !emptyLogEntry.isEmpty() {
		t.Errorf("Const emptyLogEntry should be empty")
	}
}

func TestEofLogEntry(t *testing.T) {
	if !eofLogEntry.IsEOF() {
		t.Errorf("Const eofLogEntry should be eof")
	}
}

func TestEmptyDateInLogEntry(t *testing.T) {
	entry := LogEntry{Date: time.Time{}}
	if !entry.IsEmptyDate() {
		t.Errorf("New LogEntry should have empty date")
	}
}
