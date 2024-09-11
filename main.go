package main

import (
	"errors"
	"log"
	"net/http"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func main() {
	http.Handle("/", createHandler(indexPage()))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Error:", err)
	}
}

func createHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = Page(title, r.URL.Path, body).Render(w);
	}
}

func indexPage() (string, g.Node) {
	return "Welcome!", Div(
		H1(g.Text("Welcome to the homepage!")),
		P(g.Text("This is the homepage. Please enjoy your stay.")),
	)
}

func Page(title, path string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: title,
		Language: "en",
		Head: []g.Node{
			Link(g.Attr("rel", "stylesheet"), g.Attr("href", "https://unpkg.com/mvp.css")),
			// Script(Src("https://cdn.tailwindcss.com?plugins=typography")),
		},
		Body: []g.Node{
			Container(
				body,
			),
		},
	})
}

func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}
