package main

import "fmt"
import greet "github.com/cstml/go-hello-world/greeting"

const (
  Greet = "Hey World!"
)

func main () {
  fmt.Println(greet.Greet, "!")
  fmt.Println(Greet, "!")

}
