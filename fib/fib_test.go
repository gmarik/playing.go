// http://golang.org/doc/code.html

package fib

import "testing"

func TestFib(t *testing.T) {
  if Fib(3) != (Fib(2) + Fib(1)) {
    t.Errorf("Error")
  }

  if Fib(3) != 2 {
    t.Errorf("Error: 2 != %d", Fib(3))
  }

  if Fib(5) != 5 {
    t.Errorf("Error: 5 != %d", Fib(5))
  }

  if Fib(6) != 8 {
    t.Errorf("Error: 8 != %d", Fib(5))
  }
}
