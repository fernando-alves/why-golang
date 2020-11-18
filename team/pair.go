package team

import (
	"reflect"
)

//Pair which authored a commit
type Pair interface {
	Equals(other Pair) bool
}

type pair struct {
	members *map[string]bool
}

//Equals compares a pair with other
func (p pair) Equals(other Pair) bool {
	return reflect.DeepEqual(p, other)
}

//PairWith builds pair with given members
func PairWith(members ...string) Pair {
	pairMembers := make(map[string]bool)

	for _, member := range members {
		pairMembers[member] = true
	}

	return pair{
		members: &pairMembers,
	}
}
