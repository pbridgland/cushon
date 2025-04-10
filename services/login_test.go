package services

import (
	"cushon/mocks"
	"cushon/types"
	"errors"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestLoginService_Login(t *testing.T) {
	mockRepo := mocks.DataRepo{}
	service, err := NewLoginService(&mockRepo)
	if err != nil {
		t.Fatalf("failed to create login service: %v", err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("validPass"), bcrypt.DefaultCost)
	dbErr := errors.New("db error")

	tests := []struct {
		name        string
		credentials types.Credentials
		mockSetup   func()
		expectedID  int
		expectedErr error
	}{
		{
			name: "Successful login",
			credentials: types.Credentials{
				Username: "johndoe",
				Password: "validPass",
			},
			mockSetup: func() {
				mockRepo.NextUsersWithUsernameErr = nil
				mockRepo.NextUsersWithUsernameResult = []types.User{
					{
						ID:             1,
						Username:       "johndoe",
						HashedPassword: string(hashedPassword),
					},
				}
			},
			expectedID:  1,
			expectedErr: nil,
		},
		{
			name: "Multiple users with same username",
			credentials: types.Credentials{
				Username: "duplicate",
				Password: "whatever",
			},
			mockSetup: func() {
				mockRepo.NextUsersWithUsernameErr = nil
				mockRepo.NextUsersWithUsernameResult = []types.User{
					{ID: 1, Username: "duplicate"},
					{ID: 2, Username: "duplicate"},
				}
			},
			expectedID:  0,
			expectedErr: errMultipleUsersWithSameUsername,
		},
		{
			name: "User not found",
			credentials: types.Credentials{
				Username: "nouser",
				Password: "irrelevant",
			},
			mockSetup: func() {
				mockRepo.NextUsersWithUsernameErr = nil
				mockRepo.NextUsersWithUsernameResult = []types.User{}
			},
			expectedID:  0,
			expectedErr: errMultipleUsersWithSameUsername,
		},
		{
			name: "Password mismatch",
			credentials: types.Credentials{
				Username: "johndoe",
				Password: "wrongPass",
			},
			mockSetup: func() {
				mockRepo.NextUsersWithUsernameErr = nil
				mockRepo.NextUsersWithUsernameResult = []types.User{
					{
						ID:             1,
						Username:       "johndoe",
						HashedPassword: string(hashedPassword),
					},
				}
			},
			expectedID:  0,
			expectedErr: bcrypt.ErrMismatchedHashAndPassword,
		},
		{
			name: "Repo returns error",
			credentials: types.Credentials{
				Username: "johndoe",
				Password: "irrelevant",
			},
			mockSetup: func() {
				mockRepo.NextUsersWithUsernameErr = dbErr
			},
			expectedID:  0,
			expectedErr: dbErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			id, err := service.Login(tt.credentials)
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected error: %v, got: %v", tt.expectedErr, err)
			}

			if tt.expectedErr != nil {
				return
			}

			if id != tt.expectedID {
				t.Errorf("expected ID: %d, got: %d", tt.expectedID, id)
			}
		})
	}
}
