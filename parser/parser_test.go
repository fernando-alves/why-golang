package parser

import (
	"poth/team"
	"reflect"
	"testing"
)

func TestShouldReadCommitMessage(t *testing.T) {
	commitMessage := "[TH-00][aang] Such a great commit message"

	expectedCommit := Commit{
		Authors: team.PairWith("aang"),
		Card:    "TH-00",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}

func TestShouldReadCommitWithMultipleAuthors(t *testing.T) {
	commitMessage := "[TH-00][aang|katara] Such a great commit message"

	expectedCommit := Commit{
		Authors: team.PairWith("aang", "katara"),
		Card:    "TH-00",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}

func TestShouldReadRegularCardNumber(t *testing.T) {
	commitMessage := "[TH-1234][aang] Such a great commit message"

	expectedCommit := Commit{
		Authors: team.PairWith("aang"),
		Card:    "TH-1234",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}
