/**
 * @Author: yy
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2023/06/08 13:33
 */

package demo

type UserService struct {
	repository *UserRepository
}

func NewUserService() *UserService {
	repository := NewUserRepository()
	return &UserService{repository: repository}
}

func (s *UserService) GetUsers() []UserModel {
	ids := s.repository.GetUserIds()
	return s.repository.GetUserByIds(ids)
}
