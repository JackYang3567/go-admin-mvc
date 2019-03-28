package common

import (
	. "admin-mvc/app/utils"
	"admin-mvc/app/models"	
	"encoding/json"
	//"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// GET /err?msg=
// shows the error message page
func Err(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vals := r.URL.Query()
	_, err :=  Session(w, r)
	if err != nil {
		 GenerateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		 GenerateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func Index(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
	threads, err := models.Threads()	
	//for ix, value := range threads {
	//	Info( ix,"->",template.HTMLEscape(value.Topic) )
    //}	
	//defer rows.Close()
	
		/*
	threads := [] models.Thread{} 

   
	
	for _, item := range rows{
		thread := models.Thread{}
		var topic string = item.Topic
        thread.Id  = item.Id 
		thread.Uuid = item.Uuid
		thread.Topic = topic
		thread.Topic = topic//template.HTMLEscape(w,  []byte(thread.Topic)) //template.HTMLEscapeString(item.Topic)
		thread.UserId = item.UserId
        thread.CreatedAt = item.CreatedAt
 
        threads = append(threads, thread)
	}
	*/
	if err != nil {
		 Error_message(w, r, "Cannot get threads")
	} else {
		_, err :=  Session(w, r)
		if err != nil {
			 GenerateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			 GenerateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

//api resonse json
func ThreadList(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
	threads, err := models.Threads()	
	if err != nil {
		Error_message(w, r, "Cannot get threads")
	}
	
	_data := JsonData{true, "", threads }
	data, err := json.Marshal(_data)

    if err != nil {
		Error_message(w, r, "Cannot get threads")
	}
	
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}
