package Database

import (
	"BISBackend/Config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id      int    `gorm:"primaryKey,unique,autoIncrement"`
	Secret  string `gorm:"unique"`
	Name    string
	Image   string
	Address string
}

type Message struct {
	Id     int `gorm:"primaryKey,unique,autoIncrement"`
	Text   string
	UserId int
	ChatId int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserId;references:Id" json:"-"`
	Chat   Chat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChatId;references:Id" json:"-"`
}

type Chat struct {
	Id           int `gorm:"primaryKey,unique,autoIncrement"`
	FirstUserId  int
	SecondUserId int
	FirstUser    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:FirstUserId;references:Id" json:"-"`
	SecondUser   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:SecondUserId;references:Id" json:"-"`
}

func NewDatabase() *gorm.DB {
	filePath := Config.Read().DataBasePath
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	err = db.AutoMigrate(&User{}, &Chat{}, &Message{})
	if err != nil {
		log.Fatal("Database don`t created!!!!!!")
	}
	return db
}
