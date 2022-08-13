package blackjack

var cards map[string]int

func init() {
	cards = make(map[string]int, 0)
	cards["other"] = 0
	cards["one"] = 1
	cards["two"] = 2
	cards["three"] = 3
	cards["four"] = 4
	cards["five"] = 5
	cards["six"] = 6
	cards["seven"] = 7
	cards["eight"] = 8
	cards["nine"] = 9
	cards["ten"] = 10
	cards["jack"] = 10
	cards["queen"] = 10
	cards["king"] = 10
	cards["ace"] = 11

}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerHandValue := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	if playerHandValue == 22 {
		return "P"
	} else if playerHandValue == 21 {
		if dealerCardValue != 11 && dealerCardValue != 10 {
			return "W"
		} else {
			return "S"
		}
	} else if playerHandValue >= 17 && playerHandValue <= 20 {
		return "S"
	} else if playerHandValue >= 12 && playerHandValue <= 16 {
		if dealerCardValue >= 7 {
			return "H"
		} else {
			return "S"
		}
	}
	return "H"
}
