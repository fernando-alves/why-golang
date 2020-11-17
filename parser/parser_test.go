package parser

import (
	"reflect"
	"testing"
)

func TestShouldReadCommitMessage(t *testing.T) {
	commitMessage := "[TH-00][ang] Such a great commit message"

	expectedCommit := Commit{
		Authors: []string{"ang"},
		Card:    "TH-00",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}

func TestShouldReadCommitWithMultipleAuthors(t *testing.T) {
	commitMessage := "[TH-00][ang|katara] Such a great commit message"

	expectedCommit := Commit{
		Authors: []string{"ang", "katara"},
		Card:    "TH-00",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}

func TestShouldReadRegularCardNumber(t *testing.T) {
	commitMessage := "[TH-1234][ang] Such a great commit message"

	expectedCommit := Commit{
		Authors: []string{"ang"},
		Card:    "TH-1234",
	}

	commit := Parse(commitMessage)

	if !reflect.DeepEqual(commit, expectedCommit) {
		t.Errorf("Expected author to be %v, got: %v.", expectedCommit, commit)
	}
}
