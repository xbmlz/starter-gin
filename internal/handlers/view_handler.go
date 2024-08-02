package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	r.GET("/home.html", d.home)
	r.GET("/login", d.login)
	r.GET("/register", d.register)
}

func (d viewHandler) home(c *gin.Context) {
	c.HTML(http.StatusOK, "views/home.html", gin.H{
		"title": "Home",
	})
}

func (d viewHandler) login(c *gin.Context) {
	c.HTML(http.StatusOK, "views/login.html", gin.H{
		"title": "Login",
	})
}

func (d viewHandler) register(c *gin.Context) {
	c.HTML(http.StatusOK, "views/register.html", gin.H{
		"title": "Register",
	})
}
