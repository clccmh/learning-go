package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("1) Start Game From File\n2) Start Default Game");
  scanner.Scan()
  choice := scanner.Text()
  fmt.Println("Choice " + choice)
}
