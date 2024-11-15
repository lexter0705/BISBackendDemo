package WebPerformer

import "gorm.io/gorm"

type DatabasePerformer struct {
	database *gorm.DB
}

func NewDatabasePerformer(database *gorm.DB) *DatabasePerformer {
	return &DatabasePerformer{database: database}
}
