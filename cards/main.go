package main

import "fmt"

func main() {
  // var card string = "Ace of Spades"
  // This line is equivalent to the previous.
  // Go will infer the type.
  card := "Ace of Spades"
  // The following line would throw an error.
  // := is only used when declaring a new variable.
  // card := "Five of Diamonds"
  card = "Five of Diamonds"

  fmt.Println(card)
}
