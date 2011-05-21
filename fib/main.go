package main

import fib "./fib"
import "fmt"

func main() { 
  fmt.Print("Enter number:")
  n := 0
  _, err := fmt.Scanln(&n)
  if err != nil {
    fmt.Println("Error:",err)
    return
  }
  fmt.Printf("Fib(%d)= %d\n",n,fib.Fib(n))
}
