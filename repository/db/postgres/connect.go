package postgres

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func DbOpen() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		viper.GetString(`db.host`),
		viper.GetString(`db.user`),
		viper.GetString(`db.pass`),
		viper.GetString(`db.name`),
		viper.GetString(`db.port`),
		"disable",
		"Asia/Tokyo",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func DbClose(db *gorm.DB) {
	d, _ := db.DB()
	err := d.Close()
	if err != nil {
		log.Fatal(err)
	}
}
