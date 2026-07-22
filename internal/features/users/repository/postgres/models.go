package users_postgres_repository

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type UserModel struct {
	ID        int
	Version   int
	Email     string
	CreatedAt time.Time
	Password  []byte
	FullName  string
}

func usersDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.Email,
			user.CreatedAt,
			user.Password,
			user.FullName,
		)
	}

	return userDomains
}
