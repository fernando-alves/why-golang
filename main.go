package main

import (
	"fmt"
	"io/ioutil"
	"poth/parser"
	"strings"
)

func fromFile(filePath string) []string {
	rawLogs, _ := ioutil.ReadFile(filePath)
	return strings.Split(string(rawLogs), "\n")
}

func main() {
	logEntries := fromFile("commit_history")
	commits := make([]parser.Commit, 0)
	rejectedEntries := make([]string, 0)
	for _, entry := range logEntries {
		if parser.CanParse(entry) {
			commits = append(commits, parser.Parse(entry))
		} else {
			rejectedEntries = append(rejectedEntries, entry)
		}
	}

	fmt.Printf("Parsed: %v\nRejected: %v\n", len(commits), len(rejectedEntries))
}
