package main

import "fmt"
import greet "github.com/cstml/go-hello-world/greeting"


// Say hello.
func main () {
  fmt.Println(greet.Greet, "!")
  c := make (chan int)
  go func () {
    Greet := "Hey World!"
    fmt.Println(Greet)
    c <- 0
  }()
  <-c
}
