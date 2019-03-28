package admin

import (
	"github.com/julienschmidt/httprouter"
	_"github.com/gomodule/redigo/redis"
	. "admin-mvc/app/utils"
	. "admin-mvc/app/utils/dbconn"
	 "admin-mvc/app/models"	
	_ "encoding/json"
	//"html/template"
	"net/http"	
	"time"
	"strings"
)


func Index(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {	
	
	GenerateHTML(w, "admin", "layouts/master", "layouts/public/navHeader","layouts/public/navLeft","layouts/public/footer","admin/index" )			
	
}

func Dashboard(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
	rc := RedisClient.Get()
	defer  rc.Close()

	Count := "count:dash" 
	
	rc.Send("INCR",Count)
	/*
	rc.Send("SADD",photoSet,uid)
	rc.Send("SADD",userLikeSet,photoId)

	rc.Flush()
	v,err := rc.Receive()
	if err!=nil {
		panic(err)
	}
	fmt.Println("INCR",v)
*/

	type resdata struct {
		admin models.Admin
		currentTime time.Time
	}
	 cookie, _ := r.Cookie("_cookie")
	 s := models.Session{Uuid: cookie.Value}
	 sess, _ := s.SessionByUuid()
	 admin, err := models.AdminByID(sess.AdminId)
	 res := &resdata{admin: admin, currentTime:time.Now()}
	 if err != nil{
		P("to Dashboard page view SessionByUuid===>>>",sess.AdminId)
	 }
	 P("to Dashboard page view admin===>>>",res.admin.UserName)
	 P("to Dashboard page view currentTime===>>>",res.currentTime.Format("2006-01-02 15:04:05"))
	 GenerateHTML(w, res, "layouts/main", "admin/dashboard" )
}

func SignIn(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {

	// GET Method admin signin page
	if r.Method == "GET" { 
		//P("to SignIn page view ===>>>")
		GenerateHTML(w, "admin", "layouts/nilLayout", "admin/signin" )
	}
	
	// POST Method admin signin auth
	if r.Method == "POST" {
      
			err := r.ParseForm()

			var isCap = false
			for _, v := range CapList {
				if strings.ToUpper(r.PostFormValue("captcha")) == v.Id {
					isCap = true   
				}
			}
	
		  if isCap == false {			
				 P("captcha mismatch>>>",strings.ToUpper(r.PostFormValue("captcha")))
			  return
	  	}

			admin, err := models.UserByUsername(r.PostFormValue("username"))
			if err != nil {
				Danger(err, "Cannot find admin")
			}

			if admin.Password == Encrypt(r.PostFormValue("password")) {
				session, err := admin.CreateSession()
				if err != nil {
					Danger(err, "Cannot create session")
				}
				cookie := http.Cookie{
					Name:     "_cookie",
					Value:    session.Uuid,
					HttpOnly: true,
				}
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/admin/", 302)
			} else {
				http.Redirect(w, r, "/admin/signin", 302)
			}	
		}
}

func SignOut(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//P("SignOut")
	cookie, _ := r.Cookie("_cookie")
	s := models.Session{Uuid: cookie.Value}
	err := s.DeleteAdminByUUID()
	if err != nil {
		Danger(err, "Clearing session failed")
	}
	http.Redirect(w, r, "/admin/signin", 302)
}

