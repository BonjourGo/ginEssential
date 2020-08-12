package dto

import "ginEssential/model"

type UserDTO struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func UserToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name:  user.Name,
		Phone: user.Phone,
	}

}
