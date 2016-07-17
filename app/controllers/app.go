package controllers

import (
    "github.com/revel/revel"
    "github.com/mmcdole/gofeed"
    "time"
)

func init() {
    /* 2016-07-17T22:12:42Z -> 2016-07-17 22:12:42 */
    revel.TemplateFuncs["datetimefmt"] = func(timestamp string) string {
        loc, _ := time.LoadLocation("Europe/Stockholm")
        t1, _ := time.ParseInLocation(time.RFC3339, timestamp, loc)
        fmtted := t1.Format(time.RFC1123)
        return fmtted
    }
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    fp := gofeed.NewParser()
    feed, _ := fp.ParseURL("https://joindiaspora.com/public/sebelino.atom")
	return c.Render(feed)
}
