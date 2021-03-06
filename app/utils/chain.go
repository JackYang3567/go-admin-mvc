package utils

import (
	_"encoding/json"
	_"errors"
	_"fmt"
		
	_"html/template"
	_"log"	
	_"os"
	_"path/filepath"
	_"strings"
	"admin-mvc/app/models"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//Admin authorized
func AuthorizedAdmin(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request , ps httprouter.Params) {
		cookie, err := r.Cookie("_cookie")
			//P("_cookie===>>>", cookie.Value)
			if err != nil {
				http.Redirect(w, r, "/admin/signin", 302)			
			} else {
				s := models.Session{Uuid: cookie.Value}
				valid, err := s.CheckAdmin()
				
				if err != nil {
					Danger(err, "Cannot check session")
					http.Redirect(w, r, "/admin/signin", 302)
				}
				if valid != true {
					Danger(err, "Session is not valid")
					http.Redirect(w, r, "/admin/signin", 302)
				}
			}
			h(w, r, ps)
	}
}

//Allow Access-Control-Allow-Origin
func AllowClientCors(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request , ps httprouter.Params) {
 
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")             //返回数据格式是json

	h(w, r, ps)
	}
}