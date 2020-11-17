package parser

import (
	"poth/team"
	"regexp"
	"strings"
)

var messagePattern = regexp.MustCompile(`^\[(TH-\d+)\]\[(.*)\](.*)`)

//Commit with its relevant data, authors and card number
type Commit struct {
	Authors team.Pair
	Card    string
}

//Parse message into a commit
func Parse(message string) Commit {
	matchingResult := matchMessage(message)
	return Commit{
		Authors: readPair(matchingResult),
		Card:    readCard(matchingResult),
	}
}

func readPair(matchingResult []string) team.Pair {
	authorsFromMessage := strings.Split(matchingResult[2], "|")
	return team.PairWith(authorsFromMessage...)
}

func readCard(matchingResult []string) string {
	return matchingResult[1]
}

func matchMessage(message string) []string {
	return messagePattern.FindAllStringSubmatch(message, -1)[0]
}
