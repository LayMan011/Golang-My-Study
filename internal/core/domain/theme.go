package domain

import (
	"fmt"
	"time"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type Theme struct {
	ID      int
	Version int

	Title           string
	Description     *string
	CreatedAt       time.Time
	Subject         string
	Rating          *float64
	AllRatings      int
	NumberOfRatings int
	NumberOfUsers   int
	Price           int

	AuthorUserID int
}

func NewTheme(
	id int,
	version int,
	title string,
	description *string,
	createdAt time.Time,
	subject string,
	rating *float64,
	allRatings int,
	numberOfRating int,
	numberOfUsers int,
	price int,
	authorUserID int,
) Theme {
	return Theme{
		ID:              id,
		Version:         version,
		Title:           title,
		Description:     description,
		CreatedAt:       createdAt,
		Subject:         subject,
		Rating:          rating,
		AllRatings:      allRatings,
		NumberOfRatings: numberOfRating,
		NumberOfUsers:   numberOfUsers,
		Price:           price,
		AuthorUserID:    authorUserID,
	}
}

func NewThemeUnitialized(
	title string,
	description *string,
	subject string,
	authorIserID int,
) Theme {
	return NewTheme(
		UninitailizedID,
		UninitailizedVersion,
		title,
		description,
		time.Now(),
		subject,
		nil,
		0,
		0,
		0,
		0,
		authorIserID,
	)
}

func (t *Theme) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"invalid 'Title' len: %d: %w",
			titleLen,
			core_errors.ErrInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"invalid 'Description' len: %d: %w",
				descriptionLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	subjectLen := len([]rune(t.Subject))
	if subjectLen < 1 || subjectLen > 100 {
		return fmt.Errorf(
			"invalid 'Subject' len: %d: %w",
			subjectLen,
			core_errors.ErrInvalidArgument,
		)
	}

	if t.Rating != nil {
		if *t.Rating < 0 || *t.Rating > 5 {
			return fmt.Errorf(
				"invalid 'Rating' meaning: %f: %w",
				*t.Rating,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type ThemePatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Subject     Nullable[string]
}

func NewThemePatch(
	title Nullable[string],
	description Nullable[string],
	subject Nullable[string],
) ThemePatch {
	return ThemePatch{
		Title:       title,
		Description: description,
		Subject:     subject,
	}
}

func (p *ThemePatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("'Title' can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Subject.Set && p.Subject.Value == nil {
		return fmt.Errorf("'Subject' can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (t *Theme) ApplyPatch(patch ThemePatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate theme patch: %w", err)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Subject.Set {
		tmp.Subject = *patch.Subject.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched theme: %w", err)
	}

	*t = tmp

	return nil
}
