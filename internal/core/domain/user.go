package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type User struct {
	ID      int `redis:"id"`
	Version int `redis:"version"`

	Email     string    `redis:"email"`
	CreatedAt time.Time `redis:"created_at"`
	Password  []byte    `redis:"password"`
	FullName  string    `redis:"full_name"`
}

func NewUser(
	id int,
	version int,
	email string,
	createdAt time.Time,
	password []byte,
	fullName string,
) User {
	return User{
		ID:        id,
		Version:   version,
		Email:     email,
		CreatedAt: createdAt,
		Password:  password,
		FullName:  fullName,
	}
}

func NewUserUninitailized(
	email string,
	createdAt time.Time,
	password []byte,
	fullName string,
) User {
	return NewUser(
		UninitailizedID,
		UninitailizedVersion,
		email,
		createdAt,
		password,
		fullName,
	)
}

func (u *User) Validate() error {
	basicEmailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !basicEmailRegex.MatchString(u.Email) {
		return fmt.Errorf(
			"invalid 'email': %w",
			core_errors.ErrInvalidArgument,
		)
	}

	passwordLength := len(u.Password)
	if passwordLength < 8 || passwordLength > 70 {
		return fmt.Errorf(
			"invalid 'password' len: %d: %w",
			passwordLength,
			core_errors.ErrInvalidArgument,
		)
	}

	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid 'fullName' len: %d: %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type UserPatch struct {
	Password Nullable[[]byte]
	FullName Nullable[string]
}

func NewUserPatch(
	password Nullable[[]byte],
	fullName Nullable[string],
) UserPatch {
	return UserPatch{
		Password: password,
		FullName: fullName,
	}
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf(
			"'fullName' can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Password.Set && p.Password.Value == nil {
		return fmt.Errorf(
			"'password' can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}

	if patch.Password.Set {
		tmp.Password = *patch.Password.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
