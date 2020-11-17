package team

import (
	"testing"
)

func TestShouldBeEqualsToSelf(t *testing.T) {
	pair := PairWith("korraa")

	if !pair.Equals(pair) {
		t.Errorf("Expected pair %v to be equal to %v", pair, pair)
	}
}

func TestShouldBeEqualsIfSingleMemberIsTheSame(t *testing.T) {
	pair := PairWith("korra")
	anotherPair := PairWith("korra")

	if !pair.Equals(anotherPair) {
		t.Errorf("Expected pair %v to be equal to %v", pair, pair)
	}
}

func TestShouldNotBeEqualsIfSingleMemberIsTheSame(t *testing.T) {
	pair := PairWith("korra")
	anotherPair := PairWith("kyoshi")

	if pair.Equals(anotherPair) {
		t.Errorf("Expected pair %v to be equal to %v", pair, pair)
	}
}

func TestShouldBeEqualsIfAllMembersAreTheSame(t *testing.T) {
	pair := PairWith("korra", "kyoshi")
	anotherPair := PairWith("korra", "kyoshi")

	if !pair.Equals(anotherPair) {
		t.Errorf("Expected pair %v to be equal to %v", pair, pair)
	}
}

func TestShouldBeEqualsIfAllMembersAreTheSameEvenInDifferentOrder(t *testing.T) {
	pair := PairWith("korra", "kyoshi")
	anotherPair := PairWith("kyoshi", "korra")

	if !pair.Equals(anotherPair) {
		t.Errorf("Expected pair %v to be equal to %v", pair, pair)
	}
}
