package film_grpc

import (
	"context"
	"errors"
	pb "github.com/blazee5/film-api/films/api/proto/v1"
	"github.com/blazee5/film-api/films/internal/grpc/film/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServer_CreateFilm(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.Film
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.Film{
				Title:       "Test",
				Description: "Test description",
				Genre:       "Test genre",
			},
		},
		{
			name:      "Empty credentials",
			input:     &pb.Film{},
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
		{
			name: "Duplicate error",
			input: &pb.Film{
				Title:       "Test",
				Description: "Test description",
				Genre:       "Test genre",
			},
			mockError: errors.New("title already use"),
			respError: "title already use",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			FilmServiceMock := mocks.NewFilmService(t)

			if tc.respError == "" || tc.mockError != nil {
				FilmServiceMock.On("CreateFilm", mock.Anything, tc.input).
					Return(int64(1), tc.mockError).Once()
			}

			s := Server{Service: FilmServiceMock}

			_, err := s.CreateFilm(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestServer_GetFilm(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.FilmRequest
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.FilmRequest{
				Id: 1,
			},
		},
		{
			name:      "Empty credentials",
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			FilmServiceMock := mocks.NewFilmService(t)

			if tc.respError == "" || tc.mockError != nil {
				FilmServiceMock.On("GetFilm", mock.Anything, tc.input).
					Return(&pb.Film{}, tc.mockError).Once()
			}

			s := Server{Service: FilmServiceMock}

			res, err := s.GetFilm(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}

			require.IsType(t, &pb.Film{}, res)
		})
	}
}

func TestServer_UpdateFilm(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.Film
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.Film{
				Title:       "Test",
				Description: "Test description",
				Genre:       "Test genre",
			},
		},
		{
			name:      "Empty credentials",
			input:     &pb.Film{},
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
		{
			name: "Duplicate error",
			input: &pb.Film{
				Title:       "Test",
				Description: "Test description",
				Genre:       "Test genre",
			},
			mockError: errors.New("title already use"),
			respError: "title already use",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			FilmServiceMock := mocks.NewFilmService(t)

			if tc.respError == "" || tc.mockError != nil {
				FilmServiceMock.On("UpdateFilm", mock.Anything, tc.input).
					Return(&pb.Film{}, tc.mockError).Once()
			}

			s := Server{Service: FilmServiceMock}

			res, err := s.UpdateFilm(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}

			require.IsType(t, &pb.Film{}, res)
		})
	}
}

func TestServer_DeleteFilm(t *testing.T) {
	cases := []struct {
		name      string
		input     *pb.FilmRequest
		respError string
		mockError error
	}{
		{
			name: "Success",
			input: &pb.FilmRequest{
				Id: 1,
			},
		},
		{
			name:      "Empty credentials",
			input:     &pb.FilmRequest{},
			mockError: errors.New("invalid credentials"),
			respError: "invalid credentials",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			FilmServiceMock := mocks.NewFilmService(t)

			if tc.respError == "" || tc.mockError != nil {
				FilmServiceMock.On("DeleteFilm", mock.Anything, tc.input).
					Return(tc.mockError).Once()
			}

			s := Server{Service: FilmServiceMock}

			res, err := s.DeleteFilm(context.Background(), tc.input)

			if tc.respError != "" {
				require.Equal(t, tc.respError, err.Error())
			} else {
				require.NoError(t, err)
			}

			require.IsType(t, &pb.SuccessResponse{}, res)
		})
	}
}
