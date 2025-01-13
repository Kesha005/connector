package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfigMysql struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DbName     string
	DbUsername string
	DBPassword string
}

type DBConfigPsql struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DbName     string
	DbUsername string
	DBPassword string
}



type Connector interface {
	ConnectByConfig() (*gorm.DB, error)
}

func ConnectToDB() (*gorm.DB, error) {

	var driver = os.Getenv("DB_DRIVER")

	switch driver {
	case "psql":
		fmt.Println("postgres")
		dbconfig := DBConfigPsql{
			DBDriver:   os.Getenv("DB_DRIVER"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DbUsername: os.Getenv("DB_USERNAME"),
		}

		DB, err := dbconfig.ConnectByConfig()
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		return DB, nil
	case "mysql":
		fmt.Println("mysql")
		dbconfig := DBConfigMysql{
			DBDriver:   os.Getenv("DB_DRIVER"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DbUsername: os.Getenv("DB_USERNAME"),
		}
		db, err := dbconfig.ConnectByConfig()
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		return db, nil
	default:
		err := errors.New("invalid database driver")

		log.Fatal(err)

		return nil, err
	}

}

func (dbconfig *DBConfigMysql) ConnectByConfig() (*gorm.DB, error) {

	host := dbconfig.DBHost + ":" + dbconfig.DBPort
	credentials := dbconfig.DbUsername + ":" + dbconfig.DBPassword

	dsn := credentials + "@tcp" + "(" + host + ")" + "/" + dbconfig.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root@tcp(127.0.0.1:3306)/laravel?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func (dbconfig *DBConfigPsql) ConnectByConfig() (*gorm.DB, error) {
	//timezone := os.Getenv("APP_TIMEZONE")

	args := fmt.Sprintf("host=%s port=%s dbname=%s user='%s' password=%s sslmode=%s",
		dbconfig.DBHost,
		dbconfig.DBPort,
		dbconfig.DbName,
		dbconfig.DbUsername,
		dbconfig.DBPassword,
		"prefer")

	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Connect(conn Connector) (*gorm.DB, error) {
	db, err := conn.ConnectByConfig()

	if err != nil {
		return nil, err

	}
	return db, nil
}
