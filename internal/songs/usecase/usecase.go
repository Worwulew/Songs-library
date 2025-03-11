package usecase

import (
	"context"
	"github.com/Worwulew/Songs-library/config"
	"github.com/Worwulew/Songs-library/internal/model"
	"github.com/Worwulew/Songs-library/internal/songs"
	"github.com/Worwulew/Songs-library/pkg/httpErrors"
	"github.com/Worwulew/Songs-library/pkg/logger"
	"github.com/Worwulew/Songs-library/pkg/utils"
	"github.com/pkg/errors"
)

// songsUC is songs useCase/service struct
type songsUC struct {
	cfg       *config.Config
	songsRepo songs.Repository
	logger    logger.Logger
}

// NewSongsUseCase songs useCase constructor
func NewSongsUseCase(cfg *config.Config, songsRepo songs.Repository, logger logger.Logger) songs.UseCase {
	return &songsUC{cfg: cfg, songsRepo: songsRepo, logger: logger}
}

func (s *songsUC) Create(ctx context.Context, song *model.Song) (*model.Song, error) {
	if err := utils.ValidateStruct(ctx, song); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "songsUC.Create.ValidateStruct"))
	}

	n, err := s.songsRepo.Create(ctx, song)
	if err != nil {
		return nil, err
	}

	return n, err
}

func (s *songsUC) Update(ctx context.Context, song *model.Song) (*model.Song, error) {
	if err := utils.ValidateStruct(ctx, song); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "songsUC.Update.ValidateStruct"))
	}

	n, err := s.songsRepo.Update(ctx, song)
	if err != nil {
		return nil, err
	}

	return n, err
}

func (s *songsUC) Delete(ctx context.Context, songID uint) error {
	err := s.songsRepo.Delete(ctx, songID)
	if err != nil {
		return err
	}

	return nil
}

func (s *songsUC) SongsByFields(ctx context.Context, title string, group string, query *utils.PaginationQuery) (*model.SongsList, error) {
	return s.songsRepo.SongsByFields(ctx, title, group, query)
}

func (s *songsUC) GetSongText(ctx context.Context, songID uint, query *utils.PaginationQuery) (*model.VersesList, error) {
	return s.songsRepo.GetSongText(ctx, songID, query)
}
