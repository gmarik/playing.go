
package main

import (
    "fmt";
    "io/ioutil";
    "http";
)

func main() {
    r, _, err := http.Get("http://twitter.com/statuses/public_timeline.json");
    if err != nil {
      fmt.Print("Errror", err.String())
      return
    }

    b, _ := ioutil.ReadAll(r.Body);

    fmt.Print(string(b))
}


// 6g *.go && 6l *.6 && ./6.out
