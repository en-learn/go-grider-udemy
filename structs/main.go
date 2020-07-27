package main

import "fmt"

type person struct {
  firstName string
  lastName string
}

func main() {
  // This is dependant on the order in which the fields
  // are listed in the struct definition. Ugh!
  // alex := person{"Alex", "Anderson"}

  alex := person{firstName: "Alex", lastName: "Anderson"}

  fmt.Println(alex)
}
