package domain

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestUserValidate(t *testing.T) {
	type testCase struct {
		name    string
		user    User
		wantErr bool
	}

	tests := []testCase{
		testCase{
			name: "email empty",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "",
				CreatedAt: time.Now(),
				Password:  []byte("1"),
				FullName:  "1234",
			},
			wantErr: true,
		},
		testCase{
			name: "passwordLength < 8",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte("1"),
				FullName:  "1234",
			},
			wantErr: true,
		},
		testCase{
			name: "passwordLength = 8",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte("11111111"),
				FullName:  "1234",
			},
			wantErr: false,
		},
		testCase{
			name: "passwordLength > 70",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte(strings.Repeat("1", 71)),
				FullName:  "1234",
			},
			wantErr: true,
		},
		testCase{
			name: "passwordLength = 70",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte(strings.Repeat("1", 70)),
				FullName:  "1234",
			},
			wantErr: false,
		},
		testCase{
			name: "fullNameLength < 3",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte(strings.Repeat("1", 70)),
				FullName:  "12",
			},
			wantErr: true,
		},
		testCase{
			name: "fullNameLength = 3",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte(strings.Repeat("1", 70)),
				FullName:  "123",
			},
			wantErr: false,
		},
		testCase{
			name: "fullNameLength > 100",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte(strings.Repeat("1", 70)),
				FullName:  strings.Repeat("1", 101),
			},
			wantErr: true,
		},
		testCase{
			name: "fullNameLength = 100",
			user: User{
				ID:        1,
				Version:   1,
				Email:     "user@example.com",
				CreatedAt: time.Now(),
				Password:  []byte(strings.Repeat("1", 70)),
				FullName:  strings.Repeat("1", 100),
			},
			wantErr: false,
		},
		testCase{
			name:    "return nil",
			user:    NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "123"),
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.user.Validate()
			if test.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestUserPatchValidate(t *testing.T) {
	type testCase struct {
		name    string
		patch   UserPatch
		wantErr bool
	}

	tests := []testCase{
		testCase{
			name: "non fullname",
			patch: UserPatch{
				FullName: Nullable[string]{
					Set:   true,
					Value: nil,
				},
			},
			wantErr: true,
		},
		testCase{
			name: "non password",
			patch: UserPatch{
				Password: Nullable[[]byte]{
					Set:   true,
					Value: nil,
				},
			},
			wantErr: true,
		},
		testCase{
			name:    "return nil",
			patch:   NewUserPatch(Nullable[[]byte]{Set: false, Value: nil}, Nullable[string]{Set: false, Value: nil}),
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.patch.Validate()
			if test.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestUserApplyPatch(t *testing.T) {
	type testCase struct {
		name    string
		user    User
		patch   UserPatch
		wantErr bool
	}

	passwordValid := []byte("124567890")
	passwordNonValid := []byte("123")
	fullname := "1234"

	tests := []testCase{
		testCase{
			name: "non fullname",
			user: NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "123"),
			patch: UserPatch{
				FullName: Nullable[string]{
					Set:   true,
					Value: nil,
				},
			},
			wantErr: true,
		},
		testCase{
			name:    "second validate",
			user:    NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "123"),
			patch:   NewUserPatch(Nullable[[]byte]{Set: true, Value: &passwordValid}, Nullable[string]{Set: true, Value: &fullname}),
			wantErr: false,
		},
		testCase{
			name:    "return nil",
			user:    NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "123"),
			patch:   NewUserPatch(Nullable[[]byte]{Set: true, Value: &passwordNonValid}, Nullable[string]{Set: false, Value: &fullname}),
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			passwordBefore := append([]byte(nil), test.user.Password...)
			fullNameBefore := test.user.FullName

			err := test.user.ApplyPatch(test.patch)
			if test.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !bytes.Equal(test.user.Password, passwordBefore) || test.user.FullName != fullNameBefore {
					t.Error("user changed after failed ApplyPatch()")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			passwordChanged := bytes.Equal(test.user.Password, passwordBefore)
			fullNameChanged := test.user.FullName == fullNameBefore

			if passwordChanged && fullNameChanged {
				t.Error("expected at least one field to change, but none changed")
			}
		})
	}
}
