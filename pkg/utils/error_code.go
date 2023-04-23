package utils

func ResponseCode() (response map[string]map[string]int) {
	response = map[string]map[string]int{
		"api": {
			"success":        200,
			"parameter_fail": 1001,
			"cannot_insert":  2001,
			"cannot_update":  2002,
			"cannot_delete":  2003,
			"data_not_found": 2004,
		},
	}
	return
}

func ResponseMessage() (response map[string]map[string]map[string]string) {
	response = map[string]map[string]map[string]string{
		"api": {
			"success": {
				"th": "สำเร็จ",
				"en": "Success",
			},
			"parameter_fail": {
				"th": "ข้อมูลที่ส่งมาไม่ครบถ้วน",
				"en": "incomplete information",
			},
			"cannot_insert": {
				"th": "ไม่สามารถเพิ่มข้อมูลได้",
				"en": "can't insert value !",
			},
			"cannot_update": {
				"th": "ไม่สามารถแก้ไขข้อมูลได้",
				"en": "can't update value !",
			},
			"cannot_delete": {
				"th": "ไม่สามารถลบข้อมูลได้",
				"en": "can't delete value !",
			},
			"data_not_found": {
				"th": "ไม่พบข้อมูล",
				"en": "data not found",
			},
		},
	}
	return
}
