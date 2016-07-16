package controllers

import (
    "fmt"
    "github.com/revel/revel"
    "github.com/mmcdole/gofeed"
)

type Feed struct {
	*revel.Controller
}

func (c Feed) Atom() revel.Result {
    fp := gofeed.NewParser()
    feed, _ := fp.ParseURL("https://joindiaspora.com/public/sebelino.atom")
    items := feed.Items
    for _, item := range items {
        fmt.Println(item.Title)
    }
    fmt.Println(feed.Title)
	return c.Render()
}
