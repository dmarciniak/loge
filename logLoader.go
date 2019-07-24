package loge

import (
	"bufio"
	"log"
	"github.com/dmarciniak/loge/datextractor"
	"os"
)

// LogLoader takes array of logs files and
// return LogEntry chan with sorted logs
func LogLoader(filenames []string) <-chan LogEntry {
	output := make(chan LogEntry, 100)
	go loadLogAsync(output, filenames)
	return output
}

func loadLogAsync(output chan<- LogEntry, filenames []string) {
	logEntries := make([]chan LogEntry, len(filenames))
	for fileID := range filenames {
		logEntries[fileID] = make(chan LogEntry, 20)
	}
	backed := newBucket(len(filenames))

	for fileID, filename := range filenames {
		go loadLogsFromFile(fileID, filename, logEntries[fileID])
	}

	for !backed.isClosed() {
		for fileID := backed.getIndexForFreeElement(); backed.getIndexForFreeElement() != -1; fileID = backed.getIndexForFreeElement() {
			backed[fileID] = <-logEntries[fileID]
		}

		if firstEntry := backed.popFirstEntry(); !firstEntry.isEmpty() {
			output <- firstEntry
		}
	}

	output <- eofLogEntry
}

func loadLogsFromFile(fileID int, fileName string, logEntry chan<- LogEntry) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		rawLog := scanner.Text()
		rawDate, timestamp, _ := datextractor.Extract(rawLog)
		logEntry <- LogEntry{FileID: fileID, FileName: fileName, LineNo: lineNumber, Log: rawLog, Date: timestamp, RawDate: rawDate}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	logEntry <- eofLogEntry
}
