package main

// This package will hold the Blackjack game logic

var players int

type player struct {
	playerNo int
	kind     string //player, or dealer
	card1    card
	card2    card
	hit      bool // If the player chooses hit, we change this value to true
	stand    bool // If the player chooses stand, we change this value to true
	turn     bool // Is it their turn

}

var deck deckType

var pSlice []player

type dealer struct {
	player
}

func main() {

}

func Hit(p player) { // If the player chooses Hit, they are dealt a new card, and have the choice of Hit or Stand again
	p.hit = true
	newCard(&deck)
	p.hit = false
}

func Stand(p player) { // If the player chooses stand, their turn ends, and it's the next players turn
	p.turn = false

	for i := 0; i < len(pSlice)-1; i++ {

		if pSlice[i].playerNo == p.playerNo+1 { // It is now the turn of the next player, indicated by their player numbers
			pSlice[i].turn = true
		}
	}
}

func newCard(d *deckType) {
	for i := 0; i < len(pSlice)-1; i++ {
		if pSlice[i].hit == true {
			pSlice[i].card1 = d.instance[0] // We replace their first card with the next card in the shuffled deck
		}
	}
}

func (dealer) deal(d *deckType) {
	d.Shuffle()
	for i := 0; i < len(pSlice); i++ { // Deal everyone their first card
		pSlice[i].card1 = deck.instance[0]
		if pSlice[i].kind == "player" { // If it's the dealer, the first card will not be visible
			pSlice[i].card1.visible = true
		} else {
			pSlice[i].card1.visible = false
		}
	}

	for j := 0; j < len(pSlice); j++ { // Deal everyone their second card
		pSlice[j].card2 = deck.instance[0]
		pSlice[j].card2.visible = true
	} // The second card will be visible
}

func playersInit() {
	for i := 0; i <= players; i++ {
		p := player{playerNo: i}
		if i == players { // The last player to be initialised will be the dealer
			p.kind = "dealer"
		} else {
			p.kind = "player"
		}
		pSlice = append(pSlice, p)
	}
}
