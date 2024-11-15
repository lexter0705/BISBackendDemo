package WebPerformer

import (
	"BISBackend/Database"
	"github.com/gin-gonic/gin"
)

func (performer *DatabasePerformer) CreateUser(context *gin.Context) {
	// fill after
}

func (performer *DatabasePerformer) Login(context *gin.Context) {
	var table Database.User
	performer.database.Where("Secret=?", context.Query("Secret")).First(&table)
	context.JSON(200, &table)
}

func (performer *DatabasePerformer) GetUser(context *gin.Context) {
	var table Database.User
	id := context.Query("Id")
	performer.database.First(&table, id)
	context.JSON(200, &table)
}
