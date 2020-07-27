package integration

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)
import "github.com/joho/godotenv"

func OpenDbConnection() *gorm.DB {
	_ = godotenv.Load(".env")

	pgHost := GetEnvValue("PG_HOST")
	pgPort := GetEnvValue("PG_PORT")
	pgUser := GetEnvValue("PG_USERNAME")
	pgPassword := GetEnvValue("PG_PASSWORD")
	pgDb := GetEnvValue("PG_DB")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", pgHost, pgPort, pgUser, pgDb, pgPassword)
	db, err := gorm.Open("postgres", dbUri)

	if err != nil {
		log.Fatal(fmt.Errorf("something is broken %v", err))
	}
	return db
}

func MigrateModels() {
	db := OpenDbConnection()
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Author{})
	db.Model(&Book{}).AddForeignKey("author_id", "author(id)", "RESTRICT", "RESTRICT")

}
