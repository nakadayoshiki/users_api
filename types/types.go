package types

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql との接続用
	"github.com/joho/godotenv"
)

// User 定義
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique_index"`
	Password string `gorm:"not null"`
}

const (
	mysqlUsersUserName = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHostname = "mysql_users_hostname"
	mysqlUsersSchema   = "mysql_users_schemaname"
)

// Init - users_db接続
func Init() (db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv(mysqlUsersUserName)
	password := os.Getenv(mysqlUsersPassword)
	host := os.Getenv(mysqlUsersHostname)
	schema := os.Getenv(mysqlUsersSchema)

	dataSourceName := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, schema)
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic("failed to connect database")
	}

	return
}
