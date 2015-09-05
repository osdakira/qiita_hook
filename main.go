package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func isSaveptTag(tags *simplejson.Json) bool {
	// for i := 0; i < len(tags); i++ {
	// 	tagName := tags[i].Get("name")
	// 	if tagName == "savept" {
	// 		return true
	// 	}
	// }
	return false
}

func webhook(c web.C, w http.ResponseWriter, r *http.Request) {
	js, _ := simplejson.NewFromReader(r.Body)
	fmt.Printf("%#v", js)

	tags := js.Get("tags")
	fmt.Printf("%#v", tags)

	// if isSaveptTag(tags) {
	// 	return
	// }

	// fmt.Printf("%#v", js)
	item := js.Get("item")
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
