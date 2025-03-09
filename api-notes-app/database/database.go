package database

import (
	"fmt"
	"log"
	"os"

	"github.com/azkaainurridho514/api_notes_app/config"
	"github.com/azkaainurridho514/api_notes_app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type Dbinstance struct {
	Db *gorm.DB
}
var DB Dbinstance
func Connect() {
   p := config.Config("DB_PORT")
   dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",  config.Config("DB_USER"), config.Config("DB_PASSWORD"),config.Config("DB_HOST"), p, config.Config("DB_NAME"))
   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	   Logger: logger.Default.LogMode(logger.Info),
   })
   if err != nil {
	   log.Fatal("Failed to connect to database. \n", err)
	   os.Exit(2)
   }
   log.Println("Connected")
   db.Logger = logger.Default.LogMode(logger.Info)
   log.Println("running migrations")
   db.AutoMigrate(&model.User{}, &model.Note{})
//    db.AutoMigrate(&model.User{}, &model.Role{}, &model.Company{}, &model.Request{})
//    var count int64
//    db.Model(&model.Role{}).Count(&count)
//    if count == 0 {
// 	   users := []model.Role{
// 		   {Name: "Owner"},
// 		   {Name: "Employee"},
// 	   }
// 	   db.Create(&users)
//    }
   DB = Dbinstance{
	   Db: db,
   }
}