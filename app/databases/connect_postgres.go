package databases

import (
	"fmt"

	"github.com/jeabapps/project/app/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	// Initialize connection string.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", configs.AzureConfig().Host, configs.AzureConfig().User, configs.AzureConfig().Password, configs.AzureConfig().Database)

	// Initialize connection object using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully created connection to database")

	return db, nil
}
