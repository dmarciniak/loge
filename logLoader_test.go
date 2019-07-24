package loge

import (
	"strings"
	"testing"
)

func TestLogLoader(t *testing.T) {
	output := LogLoader([]string{"testdata/file1.log", "testdata/file2.log"})

	if entry := <-output; !strings.Contains(entry.Log, "line_1") {
		t.Errorf("Wrong order in line 1")
	}

	if entry := <-output; !strings.Contains(entry.Log, "line_2") {
		t.Errorf("Wrong order in line 2")
	}

	if entry := <-output; !strings.Contains(entry.Log, "line_3") {
		t.Errorf("Wrong order in line 3")
	}

	if entry := <-output; !strings.Contains(entry.Log, "line_4") {
		t.Errorf("Wrong order in line 4")
	}

	if entry := <-output; !strings.Contains(entry.Log, "line_5") {
		t.Errorf("Wrong order in line 5")
	}

	if entry := <-output; !entry.IsEOF() {
		t.Errorf("After readed all logs lines next entry should be EOF")
	}
}
