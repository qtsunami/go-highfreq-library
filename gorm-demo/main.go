package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User2 struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(20);not null;comment:姓名;"`
	Email    string    `gorm:"type:varchar(30);not null;comment:邮箱;"`
	Age      uint8     `gorm:"type:uint;not null;comment:年龄;"`
	Birthday time.Time `gorm:"type:datetime;not null;comment:生日;"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3305)/usualdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("________")

	//email := "emojke@gmail.com"
	//birthday := time.Now()

	//user := User2{
	//	Name:     "emojke",
	//	Email:    email,
	//	Age:      18,
	//	Birthday: birthday,
	//}

	var user []User2

	db.AutoMigrate(&User2{})

	//result := db.Create(&user)

	result := db.Find(&user)
	fmt.Println(result.Error, result.RowsAffected)
	fmt.Println(user[0].Age, user[1].Age)

}

func logger() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
