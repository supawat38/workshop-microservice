package collection

import (
	"encoding/json"
	"fmt"
)

type coll map[string]interface{}

type CollSets []coll

/* ================================== */
type collSetsCollection interface {
	WhereStr(string, string) collSetsCollection
	Groupby(string) collSetsCollection
	OrderBy() collSetsCollection
	Get() collSetsCollection
}

// WhereStr 相等的值回傳
func (c CollSets) WhereStr(compareKey, mapValue string) collSetsCollection {
	var o CollSets

	// 這層是coll
	for _, subColl := range c {
		if subColl[compareKey].(string) == mapValue {
			o = append(o, subColl)
		}
	}

	return o
}

// Groupby 指定的字串做聚集
func (c CollSets) Groupby(groupKey string) collSetsCollection {

	// 紀錄有曾出現的groupby key
	var groupDataList []string
	var o CollSets

	// 這層是coll
	for _, subColl := range c {

		// 不在groupby key的紀錄
		if StringInSlice(subColl[groupKey].(string), groupDataList) == false {
			// 先記錄在一個slice
			groupDataList = append(groupDataList, subColl[groupKey].(string))

			// 建立一個新的sub object (coll) 取代 map[string]interface{} 的 interface{}
			newColl := make(coll, 1)
			newColl["groupKey"] = subColl[groupKey].(string)

			// interface{} 被換成 []coll{}
			arr := []coll{}
			arr = append(arr, subColl)

			newColl["object"] = arr
			o = append(o, newColl)

		} else {
			// 在groupby key的紀錄
			for oKey, oValue := range o {

				if oValue["groupKey"] == subColl[groupKey].(string) {

					arr := []coll{}
					arr = append(arr, subColl)

					// 這裡不知道 interface{} 被換成 []coll{}，還在認為是interface{}，所以斷言成[]coll{}
					o[oKey]["object"] = append(o[oKey]["object"].([]coll), subColl)
				}
			}
		}
	}

	return o
}

// OrderBy TODO
func (c CollSets) OrderBy() collSetsCollection {
	fmt.Println("這裡處理orderby條件(TODO)")
	return c
}

// Get return self
func (c CollSets) Get() collSetsCollection {
	json.Marshal(&c)
	// fmt.Println(string(a))
	return c
}
