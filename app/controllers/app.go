package controllers

import (
    "github.com/revel/revel"
    "github.com/mmcdole/gofeed"
    "time"
    "net/http"
    "encoding/json"
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

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

type App struct {
	*revel.Controller
}

type Post struct {
	PostType string `json:"post_type"`
	Title string `json:"title"`
    Author struct {
        DiasporaID string `json:"diaspora_id"`
    } `json:"author"`
}

func url2post(url string) Post {
    r, httperr := http.Get(url)
    defer r.Body.Close()
    if httperr != nil {
        panic(httperr)
    }
    decoder := json.NewDecoder(r.Body)
    var p Post
    err := decoder.Decode(&p)
    if err != nil {
        panic(err)
    }
    return p
}

func (c App) Index() revel.Result {
    fp := gofeed.NewParser()
    feed, err := fp.ParseURL("https://joindiaspora.com/public/sebelino.atom")
	if err != nil {
		return c.Render()
	}
    upperBound := 5
    selectedPosts := make([]*gofeed.Item, 0, upperBound)
    for _, element := range feed.Items {
        jsonurl := element.Link+".json"
        post := url2post(jsonurl)
        if post.Author.DiasporaID == "sebelino@joindiaspora.com" {
            selectedPosts = append(selectedPosts, element)
            if len(selectedPosts) == upperBound {
                break
            }
        }
    }
    feed.Items = selectedPosts
	return c.Render(feed)
}
