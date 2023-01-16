package databases

import (
	"avd.com/todo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbConn, dbOpenError := gorm.Open(sqlite.Open("./mydb.db"), &gorm.Config{})
	if dbOpenError != nil {
		panic(dbOpenError)
	}
	tableList := []interface{}{
		&models.Todo{},
	}
	if dbAutoMigrateError := dbConn.AutoMigrate(tableList...); dbAutoMigrateError != nil {
		panic(dbAutoMigrateError.Error())
	}
	DB = dbConn
}
