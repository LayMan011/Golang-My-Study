package domain

import (
	"fmt"
	"time"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type Theme struct {
	ID      int
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	Percentages int

	AuthorUserID int
}

func NewTheme(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	percentages int,
	authorUserID int,
) Theme {
	return Theme{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		CreatedAt:    createdAt,
		CompletedAt:  completedAt,
		Percentages:  percentages,
		AuthorUserID: authorUserID,
	}
}

func NewThemeUnitialized(
	title string,
	description *string,
	authorIserID int,
) Theme {
	return NewTheme(
		UninitailizedID,
		UninitailizedVersion,
		title,
		description,
		false,
		time.Now(),
		nil,
		0,
		authorIserID,
	)
}

func (t *Theme) CompletionDuration() *time.Duration {
	if !t.Completed {
		return nil
	}

	if t.CompletedAt == nil {
		return nil
	}

	duration := t.CompletedAt.Sub(t.CreatedAt)

	return &duration
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

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf(
				"'CompletedAt' can't be 'nil' if 'Completed' == true: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf(
				"'CompletedAt' can't be before 'CreatedAt': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf(
				"'CompletedAt' must be nil if 'Completed' == false: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type ThemePatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewThemePatch(
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) ThemePatch {
	return ThemePatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (p *ThemePatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("'Title' can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("'Completed' can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
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

	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value
		if tmp.Completed {
			completedAt := time.Now().UTC()
			tmp.CompletedAt = &completedAt
			tmp.Percentages = 100
		} else {
			tmp.CompletedAt = nil
			/*
				потом добавить функцию для перерасчета процентов выполнения темы
			*/
			tmp.Percentages = 0
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched theme: %w", err)
	}

	*t = tmp

	return nil
}
