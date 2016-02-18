package main

import (
  "fmt"
  "bufio"
  "os"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
  var res string;

  for res != "3" {
    printMenu()
    res = readInput()

    if res == "1" {
      fmt.Println("Hello World")
    } else if res == "2" {
      fmt.Println("I'm fine, how are you?")
    } else if res == "3" {
      fmt.Println("Quitting...")
    } else {
      fmt.Println("Invalid input")
    }

  }
}

func printMenu() {
  fmt.Println("1) Say hello")
  fmt.Println("2) Inquire Status")
  fmt.Println("3) Quit\n")
}

func readInput() string {
  scanner.Scan()
  return scanner.Text()
}
