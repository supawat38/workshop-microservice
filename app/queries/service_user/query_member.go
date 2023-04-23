package queries

import (
	"app/platform/database"

	struct_member "app/app/models/service_user"
)

// ข้อมูลสมาชิก
func GetMember(member_code string) (result []struct_member.MembersDetail) {
	var SqlWhere = ""
	if member_code != "" {
		SqlWhere += " WHERE member_code = '" + member_code + "' "
	}

	sqlStatement := ` SELECT * FROM members`
	sqlStatement += SqlWhere
	database.DBConn.Raw(sqlStatement).Scan(&result)
	return
}
