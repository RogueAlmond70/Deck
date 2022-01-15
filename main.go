package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sort"
)

type card struct {
	colour  string
	number  int
	suite   string
	royalty string
}

type deckType struct { // This currently a struct containing a slice of cards, but it itself is NOT a slice of cards. If we want to use append, it
	instance []card
} //deck is a slice of cards

var suite = []string{"spades", "diamonds", "clubs", "hearts"}
var royalty = []string{"Ace", "King", "Queen", "Jack", "Joker"}
var exSlice []int // Card numbers to remove

var jokers int
var decks int

func main() {

	cardsToRemove()
	howManyDecks()
	howManyJokers()
	newDeck()

}

func (d deckType) Shuffle() deckType { //This shuffle method takes a pointer to the deck, shuffles it, and returns it
	for i := 1; i < 52; i++ { //loop through the cards
		d.instance[i], d.instance[rand.Intn(len(d.instance)-1)] = d.instance[rand.Intn(len(d.instance)-1)], d.instance[i]
		// This takes the card at i, and the card at a random index between
		fmt.Println(&d)
	}
	return deckType{}
}

func (d deckType) Len() int {
	x := d.instance
	return len(x)
}

func (d deckType) Less(i, j int) bool {
	i = d.instance[i].number   //i is equal to the card number stored at index i
	j = d.instance[i+1].number //j is equal to the card number held at index i+1
	return i < j               //return whether the value at i is less than the value at i+1
}

func (d deckType) Swap(i, j int) {
	i = d.instance[i].number   //i is equal to the card number stored at index i
	j = d.instance[i+1].number //j is equal to the card number held at index i+1
	i, j = j, i                //swap them

}

type Interface interface{
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (d deckType) Sort(){
	sort.Slice(d.instance, func(i, j int) bool {
		return d.instance[i].number < d.instance[j].number
	})


}

func newDeck() deckType { // Have it return deckType so we can pass this as an argument to other things
	deck1 := deckType{}

	// Include logic for removing any cards

	// Include logic for using multiple decks (for things like Blackjack)
	for d := 0; d < decks; d++ {     // This outer loop repeats to compose a deck from however many decks the user wants

		for j := 0; j <= 12; j++ { // first 13 will be spades
			x := card{number: j + 1, suite: "spades", colour: "black"}

			deck1.pictureCards(x)

		}
		for k := 0; k <= 12; k++ {
			x := card{number: k + 1, suite: "diamonds", colour: "red"}

			deck1.pictureCards(x)
		}

		for l := 0; l <= 12; l++ {
			x := card{number: l + 1, suite: "clubs", colour: "black"}

			deck1.pictureCards(x)

		}

		for m := 0; m <= 12; m++ {
			x := card{number: m + 1, suite: "hearts", colour: "red"}

			deck1.pictureCards(x)

		}

		deck1.incJokers()         // Include any jokers that we need to include
		deck1.removeChosenCards() // Remove any cards we need to remove
		deck1.Sort()
	}

	fmt.Println(deck1)
	return deck1
}

func (d *deckType) removeChosenCards() {
	for i := 0; i < len(d.instance)-1; i++ { // Looping through the deck...      the first part of the for, _, would be the index

		for j := 0; j < len(exSlice); j++ {
			// ...and then for each card, looping through the "cards to remove" slice....
			if d.instance[i].number == exSlice[j] { // ... and comparing it with the current card...
				d.instance[0], d.instance[i] = d.instance[i], d.instance[0] //swap the current element with the element in position 0
				d.instance = d.instance[1:len(d.instance)]                  //then just truncate....

				//d.instance[i].suite =  "XXXXXX"}


			}
		}
	}
}







func cardsToRemove() {
	fmt.Println("Are there any card numbers you would like to exclude? Enter 'y' for yes, or 'n' for no: ")
	var text string
	fmt.Scan(&text)
	text = strings.ToLower(text)
	if text == "n" || text == strings.ToLower("no") {
		//Do nothing
	} else {
		if text == "y" || text == strings.ToLower("Yes") {
			fmt.Println("Please enter the card numbers you would like to exclude, separated by a comma, eg '2,3,7': ")
			var sequence string // creates a string variable named sequence
			fmt.Scan(&sequence)

			var splitSeq = strings.Split(sequence, ",")
			for _, value := range splitSeq {
				intValue, _ := strconv.Atoi(value)
				exSlice = append(exSlice, intValue)
			}
			fmt.Println("The cards you have chosen to remove are: ", exSlice)

		} else {
			fmt.Println("Please enter a valid input.")
			cardsToRemove()
		}
	}
}

func howManyJokers() {
	fmt.Println("How many jokers would you like? Please enter 0, 1, or 2: ") // Take input to find out how many jokers, then implement logic to include them
	var count int
	fmt.Scan(&count)
	if count > 2 || count < 0 {
		fmt.Println("Invalid number")
		howManyJokers()
	} else {

		var text string
		fmt.Printf("You would like to include %v jokers, is that correct? Enter 'y' for yes, or 'n' for no: \n", count)
		fmt.Scan(&text)

		text = strings.ToLower(text)
		if text == "y" || text == strings.ToLower("Yes") {
			fmt.Printf("You have chosen to have %d jokers", count)
			jokers = count

		} else {
			howManyJokers()
		}
	}
}


func howManyDecks() {
	fmt.Println("How many decks would you like to use? Please enter 1, 2 or 3: ") // Take input to find out how many jokers, then implement logic to include them
	var count int
	fmt.Scan(&count)
	if count > 3 || count < 1 {
		fmt.Println("Invalid number")
		howManyDecks()
	} else {

		var text string
		fmt.Printf("You would like to use %v decks, is that correct? Enter 'y' for yes, or 'n' for no: \n", count)
		fmt.Scan(&text)

		text = strings.ToLower(text)
		if text == "y" || text == strings.ToLower("Yes") {
			fmt.Printf("You have chosen to use %d decks\n", count)
			decks = count

		} else {
			howManyDecks()
		}
	}
}




func (d *deckType) incJokers() { //This method works on a pointer to a deck, and adds the jokers as required

	for i := 53; i <= 52+jokers; i++ {
		x := card{number: i, royalty: "Joker"}
		d.instance = append(d.instance, x)
	}

}

func (d *deckType) pictureCards(c card) { //This method works on a pointer to a deck, and takes a card as input
	switch c.number {
	case 1:
		c.royalty = "Ace"
	case 11:
		c.royalty = "Jack"
	case 12:
		c.royalty = "Queen"
	case 13:
		c.royalty = "King"
	}
	d.instance = append(d.instance, c)

}
