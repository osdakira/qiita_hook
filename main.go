package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func isSaveptTag(tagsJs *simplejson.Json) bool {
	fmt.Printf("%#v", tagsJs)
	tagArray, _ := tagsJs.Array()
	for i := 0; i < len(tagArray); i++ {
		tagName, _ := tagsJs.GetIndex(i).Get("name").String()
		fmt.Printf("%#v", tagName)
		if tagName == "savept" {
			return true
		}
	}
	return false
}

func webhook(c web.C, w http.ResponseWriter, r *http.Request) {
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", js)
	item := js.Get("item")

	tags := item.Get("tags")
	if !isSaveptTag(tags) {
		return
	}

	title, _ := item.Get("title").String()
	url, _ := item.Get("url").String()
	user := item.Get("user")
	// fmt.Printf("%#v", user)
	user_name, _ := user.Get("url_name").String()
	fmt.Fprintf(w, "%s, %s, %s", title, url, user_name)
}

func main() {
	goji.Post("/qiita_webhook", webhook)
	goji.Serve()
}

// http post http://localhost:8000/qiita_webhook item := []
//
// post '/qiita/webhook', provides: :json do
//   params = JSON.parse request.body.read
//   return unless params["model"] == "item"

//   title     = params["item"]["title"]
//   url       = params["item"]["url"]
//   user_name = params["item"]["user"]["url_name"]
//   yammer.notify("Qiita:\ntitle: #{title}\nposter: #{user_name}\nurl: #{url}")
// end
