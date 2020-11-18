package git

import "poth/team"

//Commit with its relevant data, authors and card number
type Commit struct {
	Authors team.Pair
	Card    string
}
