package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/micross/codingtest/models"
	"github.com/micross/codingtest/routes"
	"github.com/micross/codingtest/utils"
	"github.com/spf13/viper"
)

func setupConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/codingtest/")
	viper.AddConfigPath("$HOME/.config/codingtest")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	utils.FailOnError(err, "Failed to read config")

	env := viper.GetString("env")
	if env != models.DevelopmentMode {
		gin.SetMode(gin.ReleaseMode)
		fileName := viper.GetString("log_file")
		logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		utils.FailOnError(err, "Failed to open log file")
		gin.DefaultWriter = io.MultiWriter(logFile)
	}
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
