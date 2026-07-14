package domain

import (
	"fmt"
	"time"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type ThemeUser struct {
	ID               int
	Version          int
	Completed        bool
	AdditionAt       time.Time
	CompletedAt      *time.Time
	Percentages      int
	TotalLessons     int
	CompletedLessons int
	ThemeID          int
	UserID           int
}

func NewThemeUser(
	id int,
	version int,
	completed bool,
	additionAt time.Time,
	completeAt *time.Time,
	percentages int,
	totalLesssons int,
	compeletedLessons int,
	themeID int,
	userID int,
) ThemeUser {
	return ThemeUser{
		ID:               id,
		Version:          version,
		Completed:        completed,
		AdditionAt:       additionAt,
		CompletedAt:      completeAt,
		Percentages:      percentages,
		TotalLessons:     totalLesssons,
		CompletedLessons: compeletedLessons,
		ThemeID:          themeID,
		UserID:           userID,
	}
}

func NewThemeUserUnitialized(
	themeID int,
	userID int,
) ThemeUser {
	return NewThemeUser(
		UninitailizedID,
		UninitailizedVersion,
		false,
		time.Now(),
		nil,
		0,
		0,
		0,
		themeID,
		userID,
	)
}

func (t *ThemeUser) Validate() error {
	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf(
				"'CompletedAt' can't be 'nil' if 'Completed' == true: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		if t.CompletedAt.Before(t.AdditionAt) {
			return fmt.Errorf(
				"'CompletedAt' can't be before 'AdditionAt': %w",
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

	if t.TotalLessons < t.CompletedLessons {
		return fmt.Errorf(
			"'TotalLessons' can't be less than 'CompletedLessons': %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (t *ThemeUser) CompletionDuration() *time.Duration {
	if !t.Completed {
		return nil
	}

	if t.CompletedAt == nil {
		return nil
	}

	duration := t.CompletedAt.Sub(t.AdditionAt)

	return &duration
}

type ThemeUserPatch struct {
	Completed Nullable[bool]
}

func NewThemeUserPatch(
	completed Nullable[bool],
) ThemeUserPatch {
	return ThemeUserPatch{
		Completed: completed,
	}
}

func (p *ThemeUserPatch) Validate() error {
	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("'Completed' can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (t *ThemeUser) ApplyPatch(patch ThemeUserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate themeUser patch: %w", err)
	}

	tmp := *t

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
