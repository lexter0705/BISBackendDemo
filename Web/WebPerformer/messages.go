package WebPerformer

import (
	"BISBackend/Database"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (performer *DatabasePerformer) GetAllMessagesInChat(context *gin.Context) {
	var table []Database.Message
	id := context.Query("Id")
	performer.database.Model(&table).Where("chat_id=?", id).Find(&table)
	context.JSON(200, &table)
}

func (performer *DatabasePerformer) SendMessage(context *gin.Context) {
	chatId, err := strconv.Atoi(context.Query("ChatId"))
	if err != nil {
		context.JSON(400, gin.H{})
	}
	userId, err := strconv.Atoi(context.Query("UserId"))
	if err != nil {
		context.JSON(400, gin.H{})
	}
	text := context.Query("Text")
	message := Database.Message{Text: text, UserId: userId, ChatId: chatId}
	performer.database.Create(&message)
	context.Status(200)
}
