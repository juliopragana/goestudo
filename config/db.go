package config

import (
	"github.com/juliopragana/goestudo/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	logger := GetLogger("InitializeDB")

	dsn := "user:123456@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error connecting to dabase : %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error automigration database: %v", err)
		return nil, err
	}

	return db, nil

	// dbPath := "./db/main.db"
	// // Check if the database file exists
	// _, err := os.Stat(dbPath)
	// if os.IsNotExist(err) {
	// 	logger.Info("Database file not found, creating...")
	// 	//Create the database file and directory
	// 	err = os.MkdirAll("./db", os.ModePerm)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file, err := os.Create(dbPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file.Close()
	// }

	// //Create DB and Connect
	// db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	// if err != nil {
	// 	logger.Errorf("sqlite opening error: %v", err)
	// 	return nil, err
	// }

	// err = db.AutoMigrate(&schemas.Opening{})
	// if err != nil {
	// 	logger.Errorf("sqlite automigration error: %v", err)
	// 	return nil, err
	// }
	// // Return the DB
	// return db, nil

}
