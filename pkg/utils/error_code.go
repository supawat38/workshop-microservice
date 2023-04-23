package utils

func ResponseCode() (response map[string]map[string]int) {
	response = map[string]map[string]int{
		"api": {
			"success": 200,
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
		},
	}
	return
}
