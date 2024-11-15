package WebPerformer

import (
	"BISBackend/Database"
	"github.com/gin-gonic/gin"
)

func (performer *DatabasePerformer) GetAllMessagesInChat(context *gin.Context) {
	var table []Database.Message
	id := context.Query("Id")
	performer.database.Where("ChatId=?", id).Find(&table)
	context.JSON(200, &table)
}

func (performer *DatabasePerformer) SendMessage(context *gin.Context) {

}
