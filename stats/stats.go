package stats

import (
	"poth/git"
	"poth/team"
)

//PerPair group commits per pair
func PerPair(commits []git.Commit) map[team.Pair][]git.Commit {
	result := make(map[team.Pair][]git.Commit)
	for _, commit := range commits {
		result[commit.Authors] = append(result[commit.Authors], commit)
	}
	return result
}
