package repository

import (
	"fmt"
	er "prod/src/customlibrary/errorhandler"
	sr "prod/src/customlibrary/services"
	"time"
)

func SaveUserDetails(insertData map[string]string) int64 {

	lastId := sr.InsertPreparedQueryErr(sr.ClientDb, "user_details", insertData)
	return lastId
}

func UpdateUserData(insertData map[string]string, id string) int64 {

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	insertData["updated_date"] = now.Format("2006-01-02 15:04:05")

	condition := make(map[string]string)
	condition["id"] = fmt.Sprint(id)

	retId := sr.UpdatePrepareQuery(sr.ClientDb, "user_details", insertData, condition)
	return retId
}

func DeleteUserData(id string) int64 {

	condition := make(map[string]string)
	condition["id"] = fmt.Sprint(id)

	retId := sr.DeletePrepareQuery(sr.ClientDb, "user_details", condition)
	return retId
}

func GetAllUserData(id string) []map[string]interface{} {

	subQry := ""
	if len(id) > 0 {
		subQry = "where id=" + id + ""
	}
	query := fmt.Sprint("select id,first_name,last_name,adress,age,created_date   FROM user_details " + subQry + "")
	resultQuery := sr.SelectDirectQuery(sr.ClientDb, query)
	mapData, err := sr.ProcessResultSet(resultQuery)
	er.ErrorCheck(err)
	return mapData

}
