package database

import (
	"app/platform/logger"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	DBConn *gorm.DB
)

// Connect To Postgres
func PostgreSQLConnection(DBTableName string) (err error) {

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	//DATABASE : สำหรับเขียน
	DB_SEVER_HOST_WRITE := " host=" + os.Getenv("DB_SEVER_HOST_WRITE")
	DB_SEVER_HOST_USER_WRITE := " user=" + os.Getenv("DB_SEVER_HOST_USER_WRITE")
	DB_SEVER_HOST_PASSWORD_WRITE := " password=" + os.Getenv("DB_SEVER_HOST_PASSWORD_WRITE")
	DB_SEVER_HOST_DBNAME_WRITE := " dbname=" + DBTableName
	DB_SEVER_HOST_PORT_WRITE := " port=" + os.Getenv("DB_SEVER_HOST_PORT_WRITE")
	DB_URL_WRITE := DB_SEVER_HOST_WRITE + DB_SEVER_HOST_USER_WRITE + DB_SEVER_HOST_PASSWORD_WRITE + DB_SEVER_HOST_DBNAME_WRITE + DB_SEVER_HOST_PORT_WRITE + " sslmode=disable TimeZone=Asia/bangkok"
	dsn_write := DB_URL_WRITE

	DBConn, err = gorm.Open(postgres.Open(dsn_write), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		logger.SugarLogger.Errorf(err.Error())
		return
	}

	DBConn.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{postgres.Open(dsn_write)},
		Policy:  dbresolver.RandomPolicy{},
	}))

	sqlDB, err := DBConn.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdleConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))
	if err != nil {
		return
	}
	return
}
