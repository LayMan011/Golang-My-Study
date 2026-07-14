package users_postgres_repository

import "github.com/LayMan011/Golang-My-Study/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	Login       string
	Password    []byte
	FullName    string
	PhoneNumber *string
}

func usersDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.Login,
			user.Password,
			user.FullName,
			user.PhoneNumber,
		)
	}

	return userDomains
}
