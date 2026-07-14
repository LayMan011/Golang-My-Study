package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	Login       string
	Password    []byte
	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	login string,
	password []byte,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		Login:       login,
		Password:    password,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitailized(
	login string,
	password []byte,
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitailizedID,
		UninitailizedVersion,
		login,
		password,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	loginLength := len([]byte(u.Login))
	if loginLength < 4 || loginLength > 50 {
		return fmt.Errorf(
			"invalid 'login' len: %d: %w",
			loginLength,
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

	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf(
				"invalid 'phoneNumber' len: %d: %w",
				phoneNumberLength,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid 'fullName' len: %d: %w",
				fullNameLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type UserPatch struct {
	Password    Nullable[[]byte]
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func NewUserPatch(
	password Nullable[[]byte],
	fullName Nullable[string],
	phoneNumber Nullable[string],
) UserPatch {
	return UserPatch{
		Password:    password,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
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

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
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
