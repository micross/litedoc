package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/micross/litedoc/controllers/auth"
	"github.com/micross/litedoc/controllers/manager"
)

func Route(r *gin.Engine) {
	r.POST("/login", auth.Login)
	r.DELETE("/logout", auth.Logout)
	r.POST("/register", auth.Register)

	r.GET("/manager/users", manager.Users)
}
