package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micross/litedoc/utils"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func InitDB() {
	dsn := viper.GetString("mysql")
	db, err := gorm.Open("mysql", dsn)
	utils.FailOnError(err, "Failed to connect to MySQL")
	env := viper.GetString("env")
	if env == gin.DebugMode {
		db.LogMode(true)
	}
	DB = db
}
