package parser

import (
	"regexp"
	"strings"
)

var messagePattern = regexp.MustCompile(`^\[(TH-\d+)\]\[(.*)\](.*)`)

//Commit with its relevant data, authors and card number
type Commit struct {
	Authors []string
	Card    string
}

//Parse message into a commit
func Parse(message string) Commit {
	matchingResult := matchMessage(message)
	return Commit{
		Authors: readAuthors(matchingResult),
		Card:    readCard(matchingResult),
	}
}

func readAuthors(matchingResult []string) []string {
	authorsFromMessage := matchingResult[2]
	return strings.Split(authorsFromMessage, "|")
}

func readCard(matchingResult []string) string {
	return matchingResult[1]
}

func matchMessage(message string) []string {
	return messagePattern.FindAllStringSubmatch(message, -1)[0]
}
