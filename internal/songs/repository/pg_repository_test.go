package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Worwulew/Songs-library/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSongsRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	songsRepo := NewSongsRepository(sqlxDB)

	t.Run("Create", func(t *testing.T) {
		group := "Group Name"
		songTitle := "Song Title"
		releaseDate := "2023-01-01"
		text := "Song Lyrics"
		link := "https://example.com/song"
		createdAt := time.Now()
		updatedAt := time.Now()

		rows := sqlmock.NewRows([]string{"song_id", "group_name", "song_title", "release_date", "text", "link", "created_at", "updated_at"}).
			AddRow(1, group, songTitle, releaseDate, text, link, createdAt, updatedAt)

		song := &model.Song{
			Group:       group,
			SongTitle:   songTitle,
			ReleaseDate: releaseDate,
			Text:        text,
			Link:        link,
		}

		mock.ExpectQuery(createSong).WithArgs(song.Group, song.SongTitle, song.ReleaseDate, song.Text, song.Link).WillReturnRows(rows)

		createdSong, err := songsRepo.Create(context.Background(), song)

		require.NoError(t, err)
		require.NotNil(t, createdSong)
		require.Equal(t, song.SongTitle, createdSong.SongTitle)
		require.Equal(t, uint(1), createdSong.SongID)
		require.Equal(t, group, createdSong.Group)
		require.Equal(t, createdAt, createdSong.CreatedAt)
		require.Equal(t, updatedAt, createdSong.UpdatedAt)
	})
}

func TestSongsRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	songsRepo := NewSongsRepository(sqlxDB)

	t.Run("Update", func(t *testing.T) {
		songID := uint(1)
		group := "Updated Group Name"
		songTitle := "Updated Song Title"
		releaseDate := "2023-01-01"
		text := "Updated Song Lyrics"
		link := "https://example.com/updated-song"
		updatedAt := time.Now()

		rows := sqlmock.NewRows([]string{"song_id", "group_name", "song_title", "release_date", "text", "link", "updated_at"}).
			AddRow(songID, group, songTitle, releaseDate, text, link, updatedAt)

		song := &model.Song{
			SongID:      songID,
			Group:       group,
			SongTitle:   songTitle,
			ReleaseDate: releaseDate,
			Text:        text,
			Link:        link,
		}

		mock.ExpectQuery(updateSong).WithArgs(song.Group, song.SongTitle, song.ReleaseDate, song.Text, song.Link, song.SongID).WillReturnRows(rows)

		updatedSong, err := songsRepo.Update(context.Background(), song)

		require.NoError(t, err)
		require.NotNil(t, updatedSong)
		require.Equal(t, song.SongTitle, updatedSong.SongTitle)
		require.Equal(t, songID, updatedSong.SongID)
		require.Equal(t, group, updatedSong.Group)
		require.Equal(t, updatedAt, updatedSong.UpdatedAt)
	})
}

func TestSongsRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	songsRepo := NewSongsRepository(sqlxDB)

	t.Run("Delete", func(t *testing.T) {
		songID := uint(1)
		mock.ExpectExec(deleteSong).WithArgs(songID).WillReturnResult(sqlmock.NewResult(1, 1))

		err := songsRepo.Delete(context.Background(), songID)

		require.NoError(t, err)
	})
}
