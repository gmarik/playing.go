
package main

import (
    ts "github.com/hoisie/twitterstream"
)

func main() {
    stream := make(chan *ts.Tweet)
    client := ts.NewClient("user", "pass")
    err := client.Track([]string{ "twitter"}, stream)
    if err != nil {
        println(err.String())
    }
    for {
        tw := <- stream
        println(tw.User.Screen_name, ": ", tw.Text)
    }
}

// 6g twstream.go && 6l twstream.6 && ./6.out
