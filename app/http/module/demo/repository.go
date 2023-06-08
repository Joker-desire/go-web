/**
 * @Author: yy
 * @Description:
 * @File:  repository
 * @Version: 1.0.0
 * @Date: 2023/06/08 13:02
 */

package demo

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserIds() []int {
	return []int{1, 2}
}

func (r *UserRepository) GetUserByIds([]int) []UserModel {
	return []UserModel{
		{
			UserId: 1,
			Name:   "yy",
			Age:    18,
		},
		{
			UserId: 2,
			Name:   "yy",
			Age:    18,
		},
	}
}
