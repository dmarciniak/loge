package loge

import (
	"testing"
	"time"
)

func TestCreateNewEmpty(t *testing.T) {
	b := newBucket(3)
	if !b[0].isEmpty() || !b[1].isEmpty() || !b[2].isEmpty() {
		t.Errorf("All backet elements should be empty")
	}
}

func TestGettingIndexForFreeElement(t *testing.T) {
	b := newBucket(3)
	b[0] = LogEntry{}
	b[1] = eofLogEntry
	if b.getIndexForFreeElement() != 2 {
		t.Errorf("Only element with index 2 is empty and method should return it")
	}

	b[2] = eofLogEntry
	if b.getIndexForFreeElement() != -1 {
		t.Errorf("None of elements are free, so method should return -1")
	}
}

func TestGettingFirstEntry(t *testing.T) {
	b := newBucket(5)
	b[0] = eofLogEntry
	b[1] = LogEntry{FileID: 1}
	b[2] = emptyLogEntry
	b[3] = LogEntry{FileID: 3, Date: time.Date(2019, 01, 01, 01, 01, 01, 1002, time.Local)}
	b[4] = LogEntry{FileID: 4, Date: time.Date(2019, 01, 01, 01, 01, 01, 1001, time.Local)}

	if entry := b.popFirstEntry(); entry.FileID != 1 {
		t.Errorf("At first method should get LogEntry with empty date")
	}

	if entry := b.popFirstEntry(); entry.FileID != 4 {
		t.Errorf("Next method should get LogEntry with earlier date")
	}

	if entry := b.popFirstEntry(); entry.FileID != 3 {
		t.Errorf("After that method should get LogEntry with next date")
	}

	if entry := b.popFirstEntry(); !entry.isEmpty() {
		t.Errorf("If there are no more correctly LogEntry method should return empty LogEntry")
	}
}

func TestIsClosed(t *testing.T) {
	b := newBucket(2)
	b[0] = eofLogEntry
	b[1] = LogEntry{}
	if b.isClosed() {
		t.Errorf("Bucket isn't close because one element isn't eof")
	}

	b[1] = eofLogEntry
	if !b.isClosed() {
		t.Errorf("Bucket is close because all element are eof")
	}

}
