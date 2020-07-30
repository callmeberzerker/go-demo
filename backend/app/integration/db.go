package integration

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"

	// does side effects
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// OpenDbConnection - opens DB connection
func OpenDbConnection() *gorm.DB {
	_ = godotenv.Load(".env")

	pgHost := GetEnvValue("PG_HOST")
	pgPort := GetEnvValue("PG_PORT")
	pgUser := GetEnvValue("PG_USERNAME")
	pgPassword := GetEnvValue("PG_PASSWORD")
	pgDb := GetEnvValue("PG_DB")
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", pgHost, pgPort, pgUser, pgDb, pgPassword)
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Error().Msg(err.Error())
	}
	db.LogMode(false)
	return db
}

// MigrateModels - migrates entity models
func MigrateModels() {
	db := OpenDbConnection()
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Author{})
	db.Model(&Book{}).AddForeignKey("author_id", "author(id)", "RESTRICT", "RESTRICT")

}
