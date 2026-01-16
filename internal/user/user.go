package user

import (
	"github.com/a5932016/gin_example/databases"
	"github.com/a5932016/gin_example/internal/user/userRequest"
	"github.com/a5932016/gin_example/internal/user/userResponse"
	"github.com/a5932016/gin_example/model"
)

func List(request userRequest.List) (userResponse.List, error) {
	var response userResponse.List

	response = make(userResponse.List, 0)

	var users []model.User
	err := databases.DB.Find(&users).Error
	if err != nil {
		return response, err
	}

	for _, user := range users {
		userItem := userResponse.UserItem{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		response = append(response, userItem)
	}

	return response, nil
}
