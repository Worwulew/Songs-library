package songs

import (
	"context"
	"github.com/Worwulew/Songs-library/internal/model"
	"github.com/Worwulew/Songs-library/pkg/utils"
)

//go:generate mockgen -source=internal/songs/pg_repository.go -destination=internal/songs/mock/mock_song_pg_repository.go -package=mock

// Repository songs repository
type Repository interface {
	Create(ctx context.Context, song *model.Song) (*model.Song, error)
	Update(ctx context.Context, song *model.Song) (*model.Song, error)
	Delete(ctx context.Context, songID uint) error
	SongsByFields(ctx context.Context, title string, group string, query *utils.PaginationQuery) (*model.SongsList, error)
	GetSongText(ctx context.Context, songID uint, query *utils.PaginationQuery) (*model.VersesList, error)
}
