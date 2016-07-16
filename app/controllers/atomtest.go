package main

import (
    "fmt"
    "github.com/mmcdole/gofeed"
)

func main() {
    fmt.Println("hello")
    fp := gofeed.NewParser()
    feed, _ := fp.ParseURL("https://joindiaspora.com/public/sebelino.atom")
    items := feed.Items
    for _, item := range items {
        fmt.Println(item.Title)
    }
    fmt.Println(feed.Title)
}
