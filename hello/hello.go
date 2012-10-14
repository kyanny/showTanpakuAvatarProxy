package hello

import (
	"fmt"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type Person struct {
	Nickname string
	Username string
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, inputForm)
}

const inputForm = `
	<html>
	    <head><title>the bbs</title></head>
	    <body>
	        <form action="/post" method="post">
	            <div>nickname : <input name="nickname" type="text"></input></div>
	            <div>username : <input name="username" type="text"></input></div>
	            <div><input type="submit" value="post"></input</div>
	        </form>
	    </body>
	</html>
	`

func get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	person := new(Person)
	key    := datastore.NewKey(c, "People", r.FormValue("nickname"), 0, nil)

	err := datastore.Get(c, key, person)
	if err != nil {
		c.Infof("Get Error")
		return
	}

	avatarUrl := fmt.Sprintf("http://tanpaku.grouptube.jp/images/users/%s/icon/s.jpg", person.Username)
	c.Infof(avatarUrl)
	http.Redirect(w, r, avatarUrl, http.StatusFound)
}

func post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	key := datastore.NewKey(c, "People", r.FormValue("nickname"), 0, nil)

	var person Person
	person.Nickname = r.FormValue("nickname")
	person.Username = r.FormValue("username")
	_, err := datastore.Put(c, key, &person)
	if err != nil {
		c.Infof("error!")
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

