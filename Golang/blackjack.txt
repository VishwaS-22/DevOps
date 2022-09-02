/*
Instructions

In this exercise we will simulate the first turn of a Blackjack game.
You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string 
such as "ace", "king", "three", "two", etc. 

Note: Commonly, aces can take the value of 1 or 11 but for simplicity we will assume that they can only take the value of 11.

Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:
->Stand (S)
->Hit (H)
->Split (P)
->Automatically win (W)

Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

->If you have a pair of aces you must always split them.
->If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win.
   If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
->If your cards sum up to a value within the range [17, 20] you should always stand.
->If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
->If your cards sum up to 11 or lower you should always hit.

1. Calculate the value of any given card.
2. Implement the decision logic for the first turn.
*/

package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.

func ParseCard(card string) int { 
	switch card{
        case "one":
    	return 1
        case "two":
    	return 2
        case "three":
    	return 3
        case "four":
    	return 4
        case "five":
    	return 5
        case "six":
    	return 6
        case "seven":
    	return 7
        case "eight":
    	return 8
        case "nine":
    	return  9
        case "king","queen","jack","ten":
    	return 10
        case "ace":
    	return 11
        default:
    	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the player and one card of the dealer.

func FirstTurn(card1, card2, dealerCard string) string {
    score:=ParseCard(card1)+ParseCard(card2)
    dealerValue:=ParseCard(dealerCard)
    switch{
        case score==22:
    		return "P"
    	case score==21 && dealerValue<10:
    		return "W"
        case score==21 && dealerValue>=10:
    		return "S"
        case score>=17 && score<=20:
    		return "S"
        case score>=12 && score<=16 && dealerValue<7:
    		return "S"
        case score>=12 && score<=16 && dealerValue>=7:
    		return "H"
        default :
    		return "H"
    }
}
