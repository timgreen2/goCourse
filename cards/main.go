package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" //the "colon equals" means the same as declaring a variable of type string
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5) // since the deal function has 2 returns, two variables are required to save the info

	// hand.print()
	// remainingDeck.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my-cards")

	// cards := newDeckFromFile("my-cards")
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
