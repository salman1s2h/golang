package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RETRIES       = 10
	RETRYDURATION = time.Second * 5
)

func GetConnection(dsn string) (*gorm.DB, error) {
	counter := 0
retry:
	counter++
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("trying to connect to the database", counter, "times->", err.Error())
		if counter <= RETRIES {
			time.Sleep(RETRYDURATION)
			goto retry
		} else {
			return db, err
		}
	}
	return db, err
}
