package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/constant"
	"github.com/xbmlz/starter-gin/internal/model"
)

type viewHandler struct {
	BaseHandler
}

func NewViewHandler() viewHandler {
	return viewHandler{}
}

func (d viewHandler) Register(r *gin.Engine) {
	r.LoadHTMLGlob("web/templates/**/*")
	r.Static("/static", "web/static")

	r.GET("/", d.main)
	r.GET("/register", d.register)
	r.GET("/login", d.login)

	r.GET("/home.html", d.home)
}

func (d viewHandler) home(c *gin.Context) {
	c.HTML(http.StatusOK, "views/home.html", gin.H{
		"title": "Home",
	})
}

func (d viewHandler) login(c *gin.Context) {
	session := sessions.Default(c)
	user_id := session.Get(constant.SessionUserKey)
	if user_id != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "views/login.html", gin.H{
		"title": "Login",
	})
}

func (d viewHandler) register(c *gin.Context) {
	c.HTML(http.StatusOK, "views/register.html", gin.H{
		"title": "Register",
	})
}

func (d viewHandler) main(c *gin.Context) {
	session := sessions.Default(c)
	user_id := session.Get(constant.SessionUserKey)
	if user_id == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := model.GetUserByID(user_id.(uint))

	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "views/main.html", gin.H{
		"title": "Main",
		"user":  user,
	})
}
