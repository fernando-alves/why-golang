package stats

import (
	"poth/git"
	"poth/team"
	"reflect"
	"testing"
)

func TestPerPairShouldGroupCommitsAuthoredByTheSamePair(t *testing.T) {
	pair := team.PairWith("soka", "toph")
	commits := []git.Commit{
		aCommitFrom(pair, "JIRA-123"),
		aCommitFrom(pair, "JIRA-456"),
	}

	expectCommitsToMatch(commits, PerPair(commits)[pair], t)
}

func TestPerPairShouldGroupCommitsAuthoredByDifferentPairs(t *testing.T) {
	commitsFromSoka := []git.Commit{
		aCommitFrom(soka, "JIRA-123"),
		aCommitFrom(soka, "JIRA-789"),
	}
	commitsFromToph := []git.Commit{
		aCommitFrom(toph, "JIRA-456"),
	}

	commitsPerPair := PerPair(append(commitsFromSoka, commitsFromToph...))

	expectCommitsToMatch(commitsPerPair[soka], commitsFromSoka, t)
	expectCommitsToMatch(commitsPerPair[toph], commitsFromToph, t)
}

var soka = team.PairWith("soka")
var toph = team.PairWith("toph")

func aCommitFrom(pair team.Pair, card string) git.Commit {
	return git.Commit{
		Authors: pair,
		Card:    card,
	}
}

func expectCommitsToMatch(expected, result []git.Commit, t *testing.T) {
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected commits %v, got %v", expected, result)
	}
}
