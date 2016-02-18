package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Type Something!");
  scanner.Scan()
  fmt.Println("You typed: " + scanner.Text());
}
