package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"jwt-authentication/mocks"
	"testing"
)

func Test_Register(t *testing.T) {
	tests := []struct {
		name      string
		data      map[string]string
		setupMock func(repo *mocks.UserRepository)
		wantErr   bool
	}{
		{
			name: "Register success",
			data: map[string]string{
				"Username": "Jaana",
				"Email":    "jaana@gmail.com",
				"Password": "password",
			},
			setupMock: func(repo *mocks.UserRepository) {
				repo.On("Create", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Register fail",
			data: map[string]string{
				"Username": "Jaana",
				"Email":    "wrongmail",
				"Password": "password",
			},
			setupMock: func(repo *mocks.UserRepository) {
				repo.On("Create", mock.Anything).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mocks.UserRepository)
			tt.setupMock(repo)

			svc := NewAuthService(repo)
			err := svc.Register(tt.data)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			repo.AssertExpectations(t)
		})
	}
}
