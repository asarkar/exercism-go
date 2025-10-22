//nolint:cyclop // don't worry about cyclomatic complexity
package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		fallthrough
	case "jack":
		fallthrough
	case "queen":
		fallthrough
	case "king":
		value = 10
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	if handValue == 22 {
		return "P"
	}
	if handValue == 21 && dealerCardValue < 10 {
		return "W"
	}
	if handValue >= 17 || (handValue >= 12 && dealerCardValue < 7) {
		return "S"
	}
	return "H"
}
