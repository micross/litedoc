package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micross/litedoc/models"
	"github.com/micross/litedoc/routes"
	"github.com/micross/litedoc/utils"
	"github.com/spf13/viper"
) 

func setupConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/litedoc/")
	viper.AddConfigPath("$HOME/.config/litedoc")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	utils.FailOnError(err, "Failed to read config")
} 
 
func main() {
	setupConfig()

	models.InitDB()
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	routes.Route(app)

	app.Run(viper.GetString("listen"))
}
