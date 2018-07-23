package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "ok",
	})
}
func Logout(c *gin.Context) {

}
func Register(c *gin.Context) {

}