package services

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"jwt-authentication/mocks"
	"jwt-authentication/models"
	"testing"
)

func Test_Create_Post(t *testing.T) {
	mockRepo := new(mocks.PostRepository)
	service := NewPostService(mockRepo)

	t.Run("success", func(t *testing.T) {
		input := map[string]string{
			"title": "My Title",
			"body":  "My Body",
		}
		authorID := 1

		// Мы не знаем точное время, поэтому используем mock.MatchedBy
		mockRepo.On("Create", mock.MatchedBy(func(p *models.Post) bool {
			return p.Title == "My Title" &&
				p.Body == "My Body" &&
				p.AuthorId == authorID &&
				p.Category == "" &&
				p.Time != ""
		})).Return(nil)

		post, err := service.Create(authorID, input)

		require.NoError(t, err)
		require.NotNil(t, post)
		require.Equal(t, "My Title", post.Title)
		require.Equal(t, "My Body", post.Body)
		require.Equal(t, authorID, post.AuthorId)
		require.Equal(t, "", post.Category)
		require.NotEmpty(t, post.Time)

		mockRepo.AssertExpectations(t)
	})
}
