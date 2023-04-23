package queries

import (
	"app/platform/database"
	"crypto/sha1"
	"encoding/hex"
	"time"

	struct_authentication "app/app/models/service_authentication"
)

// สมัครสมาชิก
func Register(Parameter struct_authentication.Members) (success bool) {

	//วันที่ปัจจุบัน
	dTimeNow := time.Now().Local()
	dTimeString := dTimeNow.Format("2006-01-02 15:04:05")

	//เข้ารหัส - รหัสผ่าน
	h := sha1.New()
	h.Write([]byte(Parameter.Password))
	Password := hex.EncodeToString(h.Sum(nil))

	//เพิ่มข้อมูลที่ตารางสมาชิก
	sqlStatement := ` INSERT INTO Members
						( username , password , name , status , created_at ) `
	sqlStatement += ` 	VALUES (
							@username ,
							@password , 
							@name , 
							1 ,
							@created_at 
						) `
	if database.DBConn.Exec(sqlStatement,
		map[string]interface{}{
			"username":   Parameter.Username,
			"password":   Password,
			"name":       Parameter.Name,
			"created_at": dTimeString,
		}).Error != nil {
		return
	}

	success = true
	return
}
