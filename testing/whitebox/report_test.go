package report

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasicReport(t *testing.T) {

	getReport = func(id string) string {
		//connect to database
		return "xxxx"
	}

	report := generateReport("12DB")
	if !strings.Contains(report, ": VERIFIED") {
		t.Errorf("failed")
	}
	fmt.Println("got report : ", report)
}
