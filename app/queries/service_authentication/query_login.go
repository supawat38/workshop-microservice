package queries

import (
	struct_authentication "app/app/models/service_authentication"
	"app/platform/database"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// เข้าสู่ระบบ
func CheckMemberLogin(Parameter struct_authentication.Req_login) (Result struct_authentication.Req_login) {

	//ถอดรหัส
	h := sha1.New()
	h.Write([]byte(Parameter.Password))
	passwordEncrypt := hex.EncodeToString(h.Sum(nil))

	var tQueryWhere = ""
	tQueryWhere += ` LOWER(username) = '` + strings.ToLower(Parameter.Username) + `' AND password = '` + passwordEncrypt + `' `

	//อัพเดทข้อมูลที่ตารางผู้ใช้
	query_result := `SELECT username FROM members WHERE `
	query_result += tQueryWhere
	database.DBConn.Raw(query_result).Scan(&Result)
	return Result
}
