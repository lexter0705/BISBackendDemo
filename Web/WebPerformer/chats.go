package WebPerformer

import (
	"BISBackend/Database"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (performer *DatabasePerformer) GetAllUserChats(context *gin.Context) {
	var table []Database.Chat
	id := context.Query("Id")
	performer.database.Find(&table, "first_user_id = ? or second_user_id = ?", id, id)
	context.JSON(200, &table)
}

func (performer *DatabasePerformer) CreateNewChat(context *gin.Context) {
	firstUserNumber, err := strconv.Atoi(context.Query("FirstUserId"))
	if err != nil {
		context.JSON(400, gin.H{})
	}
	secondUserNumber, err := strconv.Atoi(context.Query("SecondUserId"))
	if err != nil {
		context.JSON(400, gin.H{})
	}
	chat := Database.Chat{FirstUserId: firstUserNumber, SecondUserId: secondUserNumber}
	performer.database.Create(&chat)
	context.Status(200)
}
