package initializers

import (
	"fmt"
	"gorm-pagination/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	DB.AutoMigrate(&models.Person{})
}

func ConnectToDB() {
	var err error
	// host := os.Getenv("HOST")
	// user := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	// dbname := os.Getenv("DBNAME")
	// port := os.Getenv("PORT")
	// sslmode := os.Getenv("SSLMODE")
	// TimeZone := os.Getenv("TIMEZONE")

	// dsn := fmt.Sprintf("host:%s user=%s password=%s dbname=%s port=%s sslmode=%s TimezZone=%s", host, user, password, dbname, port, sslmode, TimeZone)
	dsn := "host=localhost user=postgres password=1234 dbname=gorm-pagination port=8000 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err = gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed  connecting to database")
	} else {
		fmt.Printf("Database connected successfully")
	}
}
