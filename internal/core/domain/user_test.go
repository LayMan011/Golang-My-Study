package domain

import (
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
				Email: "",
			},
			wantErr: true,
		},
		testCase{
			name: "passwordLength < 8",
			user: User{
				Email:    "user@example.com",
				Password: []byte("1"),
			},
			wantErr: true,
		},
		testCase{
			name: "fullNameLength < 3",
			user: User{
				Email:    "user@example.com",
				Password: []byte("1234567890"),
				FullName: "12",
			},
			wantErr: true,
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
		patch   UserPatch
		wantErr bool
	}

	user := NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "123")

	passwordValid := []byte("124567890")
	passwordNonValid := []byte("123")
	fullname := "1234"

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
			name:    "second validate",
			patch:   NewUserPatch(Nullable[[]byte]{Set: true, Value: &passwordValid}, Nullable[string]{Set: true, Value: &fullname}),
			wantErr: false,
		},
		testCase{
			name:    "return nil",
			patch:   NewUserPatch(Nullable[[]byte]{Set: true, Value: &passwordNonValid}, Nullable[string]{Set: false, Value: &fullname}),
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := user.ApplyPatch(test.patch)
			if test.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
