package config

import(

	"tokogolangnew/entity"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_"github.com/joho/godotenv"
)


var DB *gorm.DB
var err error


func Database() {
	fmt.Println("Welcome to Toko Golang")
	dsn := "host=localhost user=postgres password=RootAccess1108 dbname=tokogolangnew_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connecting database")
	}
	fmt.Println("database connect")
}

func Migrate() {
	DB.AutoMigrate(&entity.Cart{})
}

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
	
}
