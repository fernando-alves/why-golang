package parser

import (
	"poth/git"
	"poth/team"
	"regexp"
	"strings"
)

var messagePattern = regexp.MustCompile(`^\[(TRH-\d+)\]\[(.+)\](.*)`)

//Parse message into a commit
func Parse(message string) git.Commit {
	matchingResult := matchMessage(message)
	return git.Commit{
		Authors: readPair(matchingResult),
		Card:    readCard(matchingResult),
	}
}

//CanParse check if message can be parsed properly
func CanParse(message string) bool {
	return messagePattern.MatchString(message)
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
