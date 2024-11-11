package main

import (
  "fmt"
  "os"
  "os/user"
  "nova/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Ayo %s!, This is nova programming language \n", user.Username)
  repl.Start(os.Stdin, os.Stdout)
}
