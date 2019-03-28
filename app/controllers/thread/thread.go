package thread

import (
	"fmt"
	. "admin-mvc/app/utils"
	"admin-mvc/app/models"	
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// GET /threads/new
// Show the new thread form page
func NewThread(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err :=  Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		 GenerateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func CreateThread(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess, err :=  Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			 Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			 Danger(err, "Cannot get user from session")
		}
		
		topic := template.HTMLEscapeString(r.PostFormValue("topic"))
		r.ParseForm() 
		if len(r.Form["topic"][0])==0{
			//为空的处理
		
			Danger(err, "Topic Cannot empty!")
		}else{

		
			if _, err := user.CreateThread(topic); err != nil {
				Danger(err, "Cannot create thread")
			}
		}
		http.Redirect(w, r, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//vals := r.URL.Query()
	//uuid := vals.Get("id")
	uuid := ps.ByName("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		 Error_message(w, r, "Cannot read thread ==="+uuid)
	} else {
		_, err :=  Session(w, r)
		if err != nil {
			 GenerateHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			 GenerateHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func PostThread(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess, err :=  Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			 Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			 Danger(err, "Cannot get user from session")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		//uuid := ps.ByName("id")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			 Error_message(w, r, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			 Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read/", uuid)
		http.Redirect(w, r, url, 302)
	}
}
