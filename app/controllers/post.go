package controllers

import (
	"github.com/revel/revel"
)

type Post struct {
	*revel.Controller
}
type PostModel struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var postsDB = []PostModel{
	{ID: "hoge", Title: "title1", Body: "body1"},
	{ID: "foo", Title: "title2", Body: "body2"},
}

func (c Post) Index() revel.Result {
	return c.RenderJSON(postsDB)
}

// Show the specific posts
// path: posts/:id
func (c Post) Show() revel.Result {
	id := c.Params.Route.Get("id")
	for _, v := range postsDB {
		if v.ID == id {
			return c.RenderJSON(v)
		}
	}
	return c.RenderJSON(nil)
}
