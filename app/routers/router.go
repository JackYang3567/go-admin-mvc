package routers

import (
	_"github.com/afocus/captcha"
	"net/http"
	. "admin-mvc/app/utils"
	_"admin-mvc/app/models"	
	"github.com/julienschmidt/httprouter"
	"admin-mvc/app/controllers/user"
	"admin-mvc/app/controllers/common"
	"admin-mvc/app/controllers/admin"
	"admin-mvc/app/controllers/thread"
)



// New .
func New() http.Handler {
	router := httprouter.New()

	// handle static assets
	router.ServeFiles("/static/*filepath", http.Dir(Config.Static))
	//router.NotFound = http.FileServer(http.Dir(Config.Static))

	//captcha
	router.GET( "/captcha",GetCaptcha)
	router.GET( "/captcha-query",GetCaptchaByQuery)
	router.GET( "/qrcode",Qrcode)
	

	//http.ListenAndServe(":8085", nil)

	// Admin index
	router.GET( "/admin/", AuthorizedAdmin(admin.Index))
	router.GET( "/admin/dashboard", AuthorizedAdmin(admin.Dashboard))
	router.GET( "/admin/signin", admin.SignIn)
	router.POST("/admin/signin", admin.SignIn)
	router.GET("/admin/signout", admin.SignOut)
	

	// Front index
	router.GET("/", common.Index)
	// error
	router.GET("/err", common.Err)

	//api response json
	router.GET("/api/thread/list", AllowClientCors(common.ThreadList))

	// defined in controllers/user/user.go
	router.GET("/login", user.Login)
	router.GET("/logout", user.Logout)
	router.GET("/signup", user.Signup)
	router.POST("/signup_account", user.SignupAccount)
	router.POST("/authenticate", user.Authenticate)

	// defined in controllers/thread/thread.go
	router.GET("/thread/new", thread.NewThread)
	router.POST("/thread/create", thread.CreateThread)
	router.POST("/thread/post", thread.PostThread)
	router.GET("/thread/read/:id", thread.ReadThread)

	

	return router
}