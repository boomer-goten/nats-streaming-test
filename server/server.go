package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boomer-goten/nats-streaming-test/cache"
	"github.com/boomer-goten/nats-streaming-test/model"
)

const (
	FormatData = "text/html"
	Header     = "Content-Type"
)

func RunServer(cache *cache.Cache) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		HTMLBaseForm := `<form action="/search" method="post">
		<label for="Order">Search Order:</label>
		<input type="text" id="Order" name="Order">
		<input type="submit" value="Search">
		</form>`
		HTMLBaseForm += GenerateMapHTML(*cache.GetItems())
		fmt.Fprint(w, HTMLBaseForm)
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.Form.Get("Order")
		_, ok := (*cache.GetItems())[key]
		if ok {
			data, _ := json.MarshalIndent((*cache.GetItems())[key], "", " ")
			fmt.Fprintf(w, "%s", data)
		} else {
			fmt.Fprintf(w, "Order not found")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GenerateMapHTML(m map[string]model.Order) string {
	html := "<ul>"
	for key, _ := range m {
		html += "<li>" + "Order:" + key + "</li>"
	}
	html += "</ul>"
	return html
}
