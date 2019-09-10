package main

import (
	"github.com/ahmadkarlam-ralali/latihan_go/models"
	"github.com/ahmadkarlam-ralali/latihan_go/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func main() {
	db, err := gorm.Open("mysql", "root:@/go_training?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}

	if len(os.Args[1:]) > 0 && os.Args[1] == "migrate" {
		db.
			Set("gorm:table_options", "ENGINE=InnoDB").
			AutoMigrate(&models.Todo{}, &models.User{})
		log.Println("Migrate success")
	} else {
		r := routes.SetupRouter(db)

		r.Run(":8080")
	}

	defer db.Close()
}
