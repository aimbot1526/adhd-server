package configs

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func loadEnv() error {
	const projectDirName = "adhd-server"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return err
}

func init() {

	e := loadEnv()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	gormConfig := &gorm.Config{}
	gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	conn, err := gorm.Open(postgres.Open(dbUri), gormConfig)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	db.AutoMigrate(
		&models.User{},
		&models.UserAddress{},
		&models.UserPayment{},
		&models.ShoppingSession{},
		&models.Product{},
		&models.ProductCategory{},
		&models.ProductInventory{},
		&models.CartItem{},
		&models.Discount{},
		&models.OrderDetails{},
		&models.OrderItems{},
		&models.PaymentDetails{},
	)

}

func GetDB() *gorm.DB {
	return db
}
