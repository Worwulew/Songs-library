package repository

import (
	"context"
	"database/sql"
	"github.com/Worwulew/Songs-library/internal/model"
	"github.com/Worwulew/Songs-library/internal/songs"
	"github.com/Worwulew/Songs-library/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
)

// News Repository
type songsRepo struct {
	db *sqlx.DB
}

// NewSongsRepository is song repository constructor
func NewSongsRepository(db *sqlx.DB) songs.Repository {
	return &songsRepo{db: db}
}

func (r *songsRepo) Create(ctx context.Context, song *model.Song) (*model.Song, error) {
	var createdSong model.Song

	if err := r.db.QueryRowxContext(
		ctx,
		createSong,
		&song.Group,
		&song.SongTitle,
		&song.ReleaseDate,
		&song.Text,
		&song.Link,
	).StructScan(&createdSong); err != nil {
		return nil, errors.Wrap(err, "songsRepo.Create.QueryRowxContext")
	}

	return &createdSong, nil
}

func (r *songsRepo) Update(ctx context.Context, song *model.Song) (*model.Song, error) {
	var updatedSong model.Song
	if err := r.db.QueryRowxContext(
		ctx,
		updateSong,
		&song.Group,
		&song.SongTitle,
		&song.ReleaseDate,
		&song.Text,
		&song.Link,
		&song.SongID,
	).StructScan(&updatedSong); err != nil {
		return nil, errors.Wrap(err, "songsRepo.Create.QueryRowxContext")
	}

	return &updatedSong, nil
}

func (r *songsRepo) Delete(ctx context.Context, songID uint) error {
	result, err := r.db.ExecContext(ctx, deleteSong, songID)
	if err != nil {
		return errors.Wrap(err, "songsRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "songsRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "songsRepo.Delete.rowsAffected")
	}

	return nil
}

func (r *songsRepo) SongsByFields(ctx context.Context, title string, group string, query *utils.PaginationQuery) (*model.SongsList, error) {
	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, findByTitleAndGroupCount, title, group); err != nil {
		return nil, errors.Wrap(err, "songsRepo.SongsByFields.GetContext")
	}
	if totalCount == 0 {
		return &model.SongsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Songs:      make([]*model.Song, 0),
		}, nil
	}

	var songsList = make([]*model.Song, 0, query.GetSize())
	rows, err := r.db.QueryxContext(ctx, findByTitleAndGroup, title, group, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "songsRepo.SongsByFields.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &model.Song{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "songsRepo.SongsByFields.StructScan")
		}
		songsList = append(songsList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "songsRepo.SongsByFields.rows.Err")
	}

	return &model.SongsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Songs:      songsList,
	}, nil
}

func (r *songsRepo) GetSongText(ctx context.Context, songID uint, query *utils.PaginationQuery) (*model.VersesList, error) {
	var text string
	if err := r.db.GetContext(ctx, &text, findTextByID, songID); err != nil {
		return nil, errors.Wrap(err, "songsRepo.GetSongText.GetContext")
	}

	verses := strings.Split(text, "\n")

	totalCount := len(verses)

	if totalCount == 0 {
		return &model.VersesList{
			TotalCount: 0,
			TotalPages: 0,
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    false,
			Verses:     []string{},
		}, nil
	}

	start := query.GetOffset()
	end := start + query.GetLimit()
	if end > totalCount {
		end = totalCount
	}

	paginatedVerses := verses[start:end]

	return &model.VersesList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Verses:     paginatedVerses,
	}, nil
}
