// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type Suit string

// Enum values for Suit
const (
	SuitDiamond Suit = "DIAMOND"
	SuitClub Suit = "CLUB"
	SuitHeart Suit = "HEART"
	SuitSpade Suit = "SPADE"
)

// Values returns all known values for Suit. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Suit) Values() []Suit {
	return []Suit{
		"DIAMOND",
		"CLUB",
		"HEART",
		"SPADE",
	}
}
