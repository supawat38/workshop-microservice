package database

import service_authentication "app/app/models/service_authentication"

func Init() {
	DBConn.AutoMigrate(
		&service_authentication.Members{},
	)
}
