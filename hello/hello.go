package main

import (
  "fmt"
  "log"

  "play/greetings"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0) // don't print time, source file, and line number

  names := []string{"Jarh", "Allan", "Ashley"}
  
  message, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Println(message)
}
