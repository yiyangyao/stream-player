package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usernameCookie, err := r.Cookie("username")
	sessionIdCookie, err2 := r.Cookie("session")
	if err != nil || err2 != nil || usernameCookie == nil || sessionIdCookie == nil {
		p := &HomePage{Name: "nathan"}
		t, e := template.ParseFiles("/Users/bytedance/stream-player/src/nathan.com/video-server/templates/home.html")
		if e != nil {
			log.Printf("parsing templates home.html error: %s", e)
			return
		}

		t.Execute(w, p)
	} else if len(usernameCookie.Value) != 0 && len(sessionIdCookie.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usernameCookie, err := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err != nil || err2 != nil || usernameCookie == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")

	var p *UserPage
	if len(usernameCookie.Value) != 0 {
		p = &UserPage{
			Name: usernameCookie.Value,
		}
	} else if len(fname) != 0 {
		p = &UserPage{
			Name: fname,
		}
	}

	t, e := template.ParseFiles("/Users/bytedance/stream-player/src/nathan.com/video-server/templates/userhome.html")
	if e != nil {
		log.Printf("parsing userhome.html error: %s", e)
		return
	}

	t.Execute(w, p)
}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(&Err{
			Error:     "",
			ErrorCode: "500",
		})
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(&Err{
			Error:     "",
			ErrorCode: "500",
		})
		io.WriteString(w, string(re))
		return
	}

	request(apibody, w, r)
	defer r.Body.Close()
}

func proxyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
