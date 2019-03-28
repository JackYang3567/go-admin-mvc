package user

import (
	. "admin-mvc/app/utils"
	"admin-mvc/app/models"	
	"net/http"
	
	"github.com/julienschmidt/httprouter"
)

// Index .

// GET /login
// Show the login page
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t :=  ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

// GET /signup
// Show the signup page
func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	 GenerateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func SignupAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		 Danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		 Danger(err, "Cannot create user")
	}
	http.Redirect(w, r, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	user, err := models.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		 Danger(err, "Cannot find user")
	}
	if user.Password == Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			 Danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

// GET /logout
// Logs the user out
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		 Warning(err, "Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}