/**
 * @Author: yy
 * @Description:
 * @File:  mapper
 * @Version: 1.0.0
 * @Date: 2023/06/08 12:58
 */

package demo

import (
	demoService "github.com/Joker-desire/go-web/app/provider/demo"
)

func UserModelsToUserDTOs(models []UserModel) []UserDTO {
	ret := []UserDTO{}
	for _, model := range models {
		user := UserDTO{
			ID:   model.UserId,
			Name: model.Name,
		}
		ret = append(ret, user)
	}
	return ret
}

func StudentsToUserDTOs(students []demoService.Student) []UserDTO {
	ret := []UserDTO{}
	for _, student := range students {
		user := UserDTO{
			ID:   student.ID,
			Name: student.Name,
		}
		ret = append(ret, user)
	}
	return ret
}
