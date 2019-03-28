package admin

import (
	. "admin-mvc/app/utils"
	 "admin-mvc/app/models"	
	_ "encoding/json"
	//"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"time"
	"strings"
)


func Index(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {	
	
	GenerateHTML(w, "admin", "layouts/master", "layouts/public/navHeader","layouts/public/navLeft","layouts/public/footer","admin/index" )			
	
}

func Dashboard(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
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
			
			for _, v := range CapList {
				
				P("captcha ===h>>>",v.Id)
			}
	
		   // if IsInArray(strings.ToUpper(r.PostFormValue("captcha")),CapList) == false {
			
				P("captcha mismatch>>>",strings.ToUpper(r.PostFormValue("captcha")))
			//	return
				
		//	}

			admin, err := models.UserByUsername(r.PostFormValue("username"))
			if err != nil {
				Danger(err, "Cannot find admin")
			}

			if admin.Password == models.Encrypt(r.PostFormValue("password")) {
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

