package parser

import (
	"poth/git"
	"poth/team"
	"reflect"
	"testing"
)

func TestShouldReadCommitMessage(t *testing.T) {
	commitMessage := "[TRH-00][aang] Such a great commit message"

	expectedCommit := git.Commit{
		Authors: team.PairWith("aang"),
		Card:    "TRH-00",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}

func TestShouldReadCommitWithMultipleAuthors(t *testing.T) {
	commitMessage := "[TRH-00][aang|katara] Such a great commit message"

	expectedCommit := git.Commit{
		Authors: team.PairWith("aang", "katara"),
		Card:    "TRH-00",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}

func TestShouldReadRegularCardNumber(t *testing.T) {
	commitMessage := "[TRH-1234][aang] Such a great commit message"

	expectedCommit := git.Commit{
		Authors: team.PairWith("aang"),
		Card:    "TRH-1234",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}

func TestShouldParseMessageWithAuthorsAndCard(t *testing.T) {
	commitMessage := "[TRH-1234][aang] Such a great commit message"

	if !CanParse(commitMessage) {
		t.Errorf("Expected message %v be valid", commitMessage)
	}
}
func TestShouldNotParseMessageMissingAuthors(t *testing.T) {
	commitMessage := "[TRH-1234][] Such a great commit message"

	if CanParse(commitMessage) {
		t.Errorf("Expected message %v be invalid", commitMessage)
	}
}

func TestShouldNotReadMessageMissingCard(t *testing.T) {
	commitMessage := "[][aang] Such a great commit message"

	if CanParse(commitMessage) {
		t.Errorf("Expected message %v be invalid", commitMessage)
	}
}
