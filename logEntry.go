package loge

import (
	"time"
)

// LogEntry structure with log details
// FileID is index of file in array
// RawDate is data in raw format taken form log
type LogEntry struct {
	FileID   int
	FileName string
	LineNo   int
	Log      string
	Date     time.Time
	RawDate  string
}

var (
	emptyLogEntry = LogEntry{FileID: -1, LineNo: -1}
	eofLogEntry   = LogEntry{FileID: -1, LineNo: -2}
)

func (e LogEntry) isEmpty() bool {
	return e.LineNo == -1
}

// IsEOF should be used to check if all of logs are loaded
func (e LogEntry) IsEOF() bool {
	return e.LineNo == -2
}

// IsEmptyDate should be use to check if log have any date
func (e LogEntry) IsEmptyDate() bool {
	return e.Date.IsZero()
}
