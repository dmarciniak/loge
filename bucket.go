package loge

type bucket []LogEntry

func newBucket(size int) bucket {
	b := bucket(make([]LogEntry, size))
	for index := range b {
		b[index] = emptyLogEntry
	}
	return b
}

func (b bucket) getIndexForFreeElement() int {
	for index, entry := range b {
		if entry.isEmpty() {
			return index
		}
	}
	return -1
}

func (b bucket) popFirstEntry() LogEntry {
	firstEntry := emptyLogEntry

	for _, entry := range b {
		if !entry.isEmpty() && !entry.IsEOF() && (firstEntry.isEmpty() || firstEntry.Date.After(entry.Date)) {
			firstEntry = entry
		}
	}
	if !firstEntry.isEmpty() {
		b[firstEntry.FileID] = emptyLogEntry
	}
	return firstEntry
}

func (b bucket) isClosed() bool {
	for _, entry := range b {
		if !entry.IsEOF() {
			return false
		}
	}
	return true
}
