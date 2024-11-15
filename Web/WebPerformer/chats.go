package WebPerformer

import (
	"BISBackend/Database"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (performer *DatabasePerformer) GetAllUserChats(context *gin.Context) {
	var table []Database.Chat
	id := context.Query("Id")
	performer.database.Where("FirstUserId = ? or SecondUserId = ?", id, id).Find(&table)
	context.JSON(200, &table)
}

func (performer *DatabasePerformer) CreateNewChat(context *gin.Context) {
	firstUserId, err := strconv.Atoi(context.Query("FirstUserId"))
	if err != nil {
		context.JSON(400, gin.H{})
	}
	secondUserId, err := strconv.Atoi(context.Query("SecondUserId"))
	if err != nil {
		context.JSON(400, gin.H{})
	}
	chat := Database.Chat{FirstUserId: firstUserId, SecondUserId: secondUserId}
	performer.database.Create(&chat)
	context.Status(200)
}
