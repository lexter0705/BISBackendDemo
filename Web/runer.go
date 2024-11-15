package Web

import (
	"BISBackend/Config"
	"BISBackend/Web/WebPerformer"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createServer(db *gorm.DB) *gin.Engine {
	performer := WebPerformer.NewDatabasePerformer(db)
	server := gin.Default()
	server.POST("/CreateChat", performer.CreateNewChat)
	server.POST("/GetChats", performer.GetAllUserChats)
	server.POST("/CreateNewAccount", performer.CreateUser)
	server.POST("/Login", performer.Login)
	server.POST("/SendMessage", performer.SendMessage)
	server.POST("/GetMessagesInChat", performer.GetAllMessagesInChat)
	server.POST("/GetUser", performer.GetUser)
	return server
}

func RunServer() {
	config := Config.Read()
	db, err := gorm.Open(sqlite.Open(string(config.DataBasePath)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	server := createServer(db)
	err = server.Run(config.Port)
	if err != nil {
		panic("failed start server")
	}
}
