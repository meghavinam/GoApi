/**
 * For handling apis apis
 *
 */
package controller

import (
	rp "prod/src/customlibrary/repository"
)

/**
 * For saving the user in to database
 * @input data from api
 * @return status
 */

func SaveUserDetails(requestData map[string]string) bool {
	status := false

	if len(requestData) > 0 {
		addUser := rp.SaveUserDetails(requestData)
		if addUser > 0 {
			status = true
		}

	}
	return status
}

/**
 * For update the user details in to database
 * @input post inouts and update id from api
 * @return status
 */

func UpdateUserDetails(requestData map[string]string) bool {
	status := false

	if len(requestData) > 0 {

		updateUser := rp.UpdateUserData(requestData, requestData["id"])
		if updateUser > 0 {

			status = true
		}

	}
	return status
}

/**
 * For delete the user from database
 * @input delete id from api
 * @return status
 */

func DeleteUserDetails(requestData map[string]string) bool {
	status := false

	if len(requestData) > 0 {

		deleteUser := rp.DeleteUserData(requestData["id"])
		if deleteUser > 0 {

			status = true
		}

	}
	return status
}

/**
 * For listing all users from the database
 *
 * @return status ,user lists
 */

func GetAllUserDetails(requestData map[string]string) (bool, []map[string]interface{}) {
	status := true
	_, ok := requestData["id"]
	var id string
	if ok {
		id = requestData["id"]
	}

	lists := rp.GetAllUserData(id)
	if len(lists) > 0 {
		status = true
	}

	return status, lists
}
