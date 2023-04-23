package database

import (
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBMysqlConn *gorm.DB
)

//Connect To Mysql
func MySQLConnection(tHost string, tDbuser string, tDbpass string, tDatabaseName string) (err error) {

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	var datetimePrecision = 2
	tConnection := tDbuser + ":" + tDbpass + "@tcp(" + tHost + ":3306)/" + tDatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	DBMysqlCon, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       tConnection,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DefaultDatetimePrecision:  &datetimePrecision,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := DBMysqlCon.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdleConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	DBMysqlConn = DBMysqlCon

	if err != nil {
		return
	}

	// sqlDB.Close()

	return
}

//CloseConnectionDB
func CloseConnectDBMysql() {
	sqlDB, _ := DBMysqlConn.DB()
	sqlDB.Close()
}
