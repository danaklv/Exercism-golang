package tasks2

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        return 11
        case "queen", "king", "jack", "ten":
        return 10
        case "nine":
        return 9
        case "eight":
        return 8
        case "seven":
        return 7
        case "six":
        return 6
        case "five":
        return 5
        case "four":
        return 4
        case "three":
        return 3
        case "two":
        return 2
        default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    cardNum1 := ParseCard(card1)
    cardNum2 := ParseCard(card2)
    dealerCardNum := ParseCard(dealerCard)
	switch {
        case cardNum1 == 11 && cardNum2 == 11:
        return "P"
        case cardNum1 + cardNum2 == 21:
        if dealerCardNum >= 10 {
            return "S"
        } else {
            return "W"
        }
        case cardNum1 + cardNum2 >= 17 && cardNum1 + cardNum2 <= 20:
            return "S"
        case cardNum1 + cardNum2 >= 12 && cardNum1 + cardNum2 <= 16:
        if dealerCardNum >= 7 {
            return "H"
        }
        return "S"
        default :
        return "H"
    }
}