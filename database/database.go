package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm/logger"

	"github.com/lambaharsh01/surveyItBackend/models/databaseSchema"
	"github.com/lambaharsh01/surveyItBackend/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dataBaseInstance *gorm.DB

func InitDb() *gorm.DB {
	var err error
	dataBaseInstance, err = connectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return nil
	}

	if err := performMigrations(); err != nil {
		log.Fatalf("Could not perform migrations: %v", err)
	}

	if err := MigrateWithDefaultData(); err != nil {
		log.Fatalf("Could not migrate with default data: %v", err)
	}

	return dataBaseInstance
}

func connectDB() (*gorm.DB, error) {
	dbUsername := utils.GetEnv("DB_USERNAME")
	dbPassword := utils.GetEnv("DB_PASSWORD")
	dbName := utils.GetEnv("DB_NAME")
	dbHost := utils.GetEnv("DB_HOST")
	dbPort := utils.GetEnv("DB_PORT")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	maxIdleConns, _ := strconv.Atoi(utils.GetEnv("DB_MAX_IDLE_CONNS"))
	maxOpenConns, _ := strconv.Atoi(utils.GetEnv("DB_MAX_OPEN_CONNS"))
	connMaxLifetime, _ := strconv.Atoi(utils.GetEnv("DB_CONN_MAX_LIFETIME")) // in seconds

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	return databaseConnection, nil
}

func performMigrations() error {
	err := dataBaseInstance.AutoMigrate(
		&databaseSchema.Users{},
		&databaseSchema.QuestionType{},
		&databaseSchema.FileType{},
		&databaseSchema.SurveySchema{},
		&databaseSchema.QuestionSchema{},
		&databaseSchema.SurveyResponseSummary{},
		&databaseSchema.SurveyResponseDetails{},
	)
	return err
}

func MigrateWithDefaultData() error {

	// transaction := dataBaseInstance.Begin()

	// defer func() {
	// 	if recover := recover(); recover != nil {
	// 		transaction.Rollback()
	// 	}
	// }()

	// transaction.Exec("DELETE FROM file_types WHERE 1=1")
	// if err := transaction.Create(&databaseSchema.DefaultFileTypes).Error; err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }

	// transaction.Exec("DELETE FROM question_types WHERE 1=1")
	// if err := transaction.Create(&databaseSchema.DefaultQuestionTypes).Error; err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }

	// return transaction.Commit().Error
	return nil

}

func CloseDB() {
	sqlDB, err := dataBaseInstance.DB()
	if err != nil {
		log.Printf("Error retrieving database instance: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection pool: %v", err)
	}
}
