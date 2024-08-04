package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/constant"
)

type viewHandler struct {
	BaseHandler
}

func NewViewHandler() viewHandler {
	return viewHandler{}
}

func (d viewHandler) Register(r *gin.Engine) {
	r.LoadHTMLGlob("web/templates/**/*.html")
	r.Static("/static", "web/static")
	r.Static("/uploads", "uploads")

	r.GET("/main.html", d.main)
	r.GET("/register.html", d.register)
	r.GET("/login.html", d.login)

	r.GET("/home.html", d.home)
	r.GET("/profile.html", d.profile)
	r.GET("/password.html", d.password)
}

func (h viewHandler) home(c *gin.Context) {
	c.HTML(http.StatusOK, "views/home.html", gin.H{
		"title": "Home",
	})
}

func (h viewHandler) login(c *gin.Context) {
	session := sessions.Default(c)
	user_id := session.Get(constant.SessionUserKey)
	if user_id != nil {
		c.Redirect(http.StatusFound, "/main.html")
		return
	}

	c.HTML(http.StatusOK, "views/login.html", gin.H{
		"title": "Login",
	})
}

func (h viewHandler) register(c *gin.Context) {
	c.HTML(http.StatusOK, "views/register.html", gin.H{
		"title": "Register",
	})
}

func (h viewHandler) main(c *gin.Context) {
	c.HTML(http.StatusOK, "views/main.html", gin.H{
		"title": "Main",
		"user":  h.GetCurrentUser(c),
	})
}

func (h viewHandler) profile(c *gin.Context) {
	c.HTML(http.StatusOK, "views/profile.html", gin.H{
		"title": "Profile",
	})
}

func (h viewHandler) password(c *gin.Context) {
	c.HTML(http.StatusOK, "views/password.html", gin.H{
		"title": "Password",
	})
}
