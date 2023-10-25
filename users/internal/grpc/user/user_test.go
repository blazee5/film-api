package user_grpc

import (
	"context"
	"errors"
	pb "github.com/blazee5/film-api/api/proto/v1"
	"github.com/blazee5/film-api/users/internal/grpc/user/mocks"
	"github.com/blazee5/film-api/users/internal/models"
	"github.com/blazee5/film-api/users/lib/auth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServer_SignUp(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.User
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.User{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "password",
			},
		},
		{
			name: "Empty credentials",
			input: &pb.User{
				Name:     "",
				Email:    "",
				Password: "",
			},
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
		{
			name: "Duplicate error",
			input: &pb.User{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "password",
			},
			mockError: errors.New("email already use"),
			respError: "email already use",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			UserServiceMock := mocks.NewUserService(t)

			if tc.respError == "" || tc.mockError != nil {
				UserServiceMock.On("CreateUser", mock.Anything, tc.input).
					Return(int64(1), tc.mockError).Once()
			}

			s := Server{Service: UserServiceMock}

			_, err := s.SignUp(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestServer_SignIn(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.User
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.User{
				Email:    "test@gmail.com",
				Password: "password",
			},
		},
		{
			name: "Empty credentials",
			input: &pb.User{
				Email:    "",
				Password: "",
			},
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			UserServiceMock := mocks.NewUserService(t)

			if tc.respError == "" || tc.mockError != nil {
				UserServiceMock.On("ValidateUser", mock.Anything, tc.input.Email, auth.GenerateHashPassword(tc.input.Password)).
					Return(&models.User{Email: tc.input.Email, Password: tc.input.Password}, tc.mockError).Once()
			}

			s := Server{Service: UserServiceMock}

			_, err := s.SignIn(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}

		})
	}
}

func TestServer_GetUser(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.UserRequest
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.UserRequest{
				Id: 1,
			},
		},
		{
			name:      "Empty credentials",
			input:     &pb.UserRequest{},
			respError: "you are not this users",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			UserServiceMock := mocks.NewUserService(t)

			if tc.respError == "" || tc.mockError != nil {
				UserServiceMock.On("GetUser", mock.Anything, tc.input).
					Return(&pb.UserInfo{}, tc.mockError).Once()
			}

			s := Server{Service: UserServiceMock}

			res, err := s.GetUser(context.WithValue(context.Background(), "user_id", 1), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}

			require.IsType(t, &pb.UserInfo{}, res)
		})
	}
}

func TestServer_UpdateUser(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.User
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.User{
				Id:       1,
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "password",
			},
		},
		{
			name: "Empty credentials",
			input: &pb.User{
				Name:     "",
				Email:    "",
				Password: "",
			},
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
		{
			name: "Duplicate error",
			input: &pb.User{
				Id:       1,
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "password",
			},
			mockError: errors.New("email already use"),
			respError: "email already use",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			UserServiceMock := mocks.NewUserService(t)

			if tc.respError == "" || tc.mockError != nil {
				UserServiceMock.On("UpdateUser", mock.Anything, tc.input).
					Return(&pb.User{}, tc.mockError).Once()
			}

			s := Server{Service: UserServiceMock}

			res, err := s.UpdateUser(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}

			require.IsType(t, &pb.User{}, res)
		})
	}
}

func TestServer_DeleteUser(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.UserRequest
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.UserRequest{
				Id: 1,
			},
		},
		{
			name:      "Empty credentials",
			input:     &pb.UserRequest{},
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			UserServiceMock := mocks.NewUserService(t)

			if tc.respError == "" || tc.mockError != nil {
				UserServiceMock.On("DeleteUser", mock.Anything, tc.input).
					Return(tc.mockError).Once()
			}

			s := Server{Service: UserServiceMock}

			res, err := s.DeleteUser(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}

			require.IsType(t, &pb.SuccessResponse{}, res)
		})
	}
}
