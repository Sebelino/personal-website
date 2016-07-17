package controllers

import (
    "github.com/revel/revel"
    "github.com/mmcdole/gofeed"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    fp := gofeed.NewParser()
    feed, _ := fp.ParseURL("https://joindiaspora.com/public/sebelino.atom")
	return c.Render(feed)
}
