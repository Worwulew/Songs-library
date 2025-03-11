package http

import (
	"encoding/json"
	"github.com/Worwulew/Songs-library/internal/model"
	"github.com/Worwulew/Songs-library/internal/songs/mock"
	"github.com/Worwulew/Songs-library/pkg/converter"
	"github.com/Worwulew/Songs-library/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestSongHandler_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockSongsUC := mock.NewMockUseCase(ctrl)
	songHandlers := NewSongHandlers(nil, mockSongsUC, apiLogger)

	router := gin.Default()
	router.POST("/songs/create", songHandlers.Create)

	song := &model.Song{
		Group:       "Test Group",
		SongTitle:   "Test Song Title",
		ReleaseDate: "2023.10.01",
		Text:        "This is a test song text",
		Link:        "https://example.com/song",
	}

	buf, err := converter.AnyToBytesBuffer(song)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/songs/create", strings.NewReader(buf.String()))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	mockSongsUC.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(song, nil)

	router.ServeHTTP(res, req)

	require.Equal(t, http.StatusCreated, res.Code)

	var responseSong model.Song
	err = json.Unmarshal(res.Body.Bytes(), &responseSong)
	require.NoError(t, err)

	require.Equal(t, song.Group, responseSong.Group)
	require.Equal(t, song.SongTitle, responseSong.SongTitle)
	require.Equal(t, song.ReleaseDate, responseSong.ReleaseDate)
	require.Equal(t, song.Text, responseSong.Text)
	require.Equal(t, song.Link, responseSong.Link)
}

func TestSongHandler_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockSongsUC := mock.NewMockUseCase(ctrl)
	songHandlers := NewSongHandlers(nil, mockSongsUC, apiLogger)

	router := gin.Default()
	router.PUT("/songs/update/:id", songHandlers.Update)

	song := &model.Song{
		SongID:      1,
		Group:       "Updated Test Group",
		SongTitle:   "Updated Test Song Title",
		ReleaseDate: "2023.10.02",
		Text:        "This is an updated test song text",
		Link:        "https://example.com/updated-song",
	}

	buf, err := converter.AnyToBytesBuffer(song)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPut, "/songs/update/1", strings.NewReader(buf.String()))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	mockSongsUC.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(song, nil)

	router.ServeHTTP(res, req)

	require.Equal(t, http.StatusOK, res.Code)

	var responseSong model.Song
	err = json.Unmarshal(res.Body.Bytes(), &responseSong)
	require.NoError(t, err)

	require.Equal(t, song.SongID, responseSong.SongID)
	require.Equal(t, song.Group, responseSong.Group)
	require.Equal(t, song.SongTitle, responseSong.SongTitle)
	require.Equal(t, song.ReleaseDate, responseSong.ReleaseDate)
	require.Equal(t, song.Text, responseSong.Text)
	require.Equal(t, song.Link, responseSong.Link)
}

func TestSongHandler_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockSongsUC := mock.NewMockUseCase(ctrl)
	songHandlers := NewSongHandlers(nil, mockSongsUC, apiLogger)

	router := gin.Default()
	router.DELETE("/songs/delete/:id", songHandlers.Delete)

	songID := 1

	req := httptest.NewRequest(http.MethodDelete, "/songs/delete/"+strconv.Itoa(songID), nil)
	res := httptest.NewRecorder()

	mockSongsUC.EXPECT().
		Delete(gomock.Any(), uint(songID)).
		Return(nil)

	router.ServeHTTP(res, req)

	require.Equal(t, http.StatusOK, res.Code)

	var responseBody map[string]string
	err := json.Unmarshal(res.Body.Bytes(), &responseBody)
	require.NoError(t, err)
	require.Equal(t, "deleted", responseBody["status"])
}
