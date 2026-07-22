package users_transport_http

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type UserDTOResponse struct {
	ID        int       `json:"id" example:"10"`
	Version   int       `json:"version" example:"3"`
	Email     string    `json:"email" example:"example@gmail.com"`
	CreatedAt time.Time `json:"created_at" example:"2026-02-26T10:30:00Z"`
	Password  []byte    `json:"password"`
	FullName  string    `json:"full_name" example:"Ivan Ivanov"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:       user.ID,
		Version:  user.Version,
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
