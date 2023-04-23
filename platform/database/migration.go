package database

func Init() {
	DBConn.AutoMigrate()
}
