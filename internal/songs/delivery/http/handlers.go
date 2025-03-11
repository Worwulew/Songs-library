package http

import (
	"github.com/Worwulew/Songs-library/config"
	"github.com/Worwulew/Songs-library/internal/model"
	"github.com/Worwulew/Songs-library/internal/songs"
	"github.com/Worwulew/Songs-library/pkg/httpErrors"
	"github.com/Worwulew/Songs-library/pkg/logger"
	"github.com/Worwulew/Songs-library/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SongHandler is song handlers
type SongHandler struct {
	cfg     *config.Config
	songsUC songs.UseCase
	logger  logger.Logger
}

// NewSongHandlers News handlers constructor
func NewSongHandlers(cfg *config.Config, songsUC songs.UseCase, logger logger.Logger) *SongHandler {
	return &SongHandler{cfg: cfg, songsUC: songsUC, logger: logger}
}

// Create godoc
// @Summary Create song
// @Description Create song handler
// @Tags Song
// @Param        request   body      model.Song  true  "request for creating song"
// @Accept json
// @Produce json
// @Success 201 {object} model.Song
// @Failure 400 {object} httpErrors.RestError "Invalid request body"
// @Failure 500 {object} httpErrors.RestError "Internal server error"
// @Router /songs [post]
func (h *SongHandler) Create(c *gin.Context) {
	var song *model.Song
	var err error
	if err := c.ShouldBindJSON(&song); err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	// For external API use
	/*song, err = FetchSongDetails(song)
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}*/

	createdSong, err := h.songsUC.Create(c, song)
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, createdSong)
}

// Update godoc
// @Summary Update song
// @Description Update song handler
// @Tags Song
// @Param id path int true "song_id"
// @Param        request   body      model.Song  true  "request for updating song"
// @Accept json
// @Produce json
// @Success 200 {object} model.Song
// @Failure 400 {object} httpErrors.RestError "Invalid request body"
// @Failure 500 {object} httpErrors.RestError "Internal server error"
// @Router /songs/{id} [put]
func (h *SongHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(httpErrors.BadQueryParams))
		return
	}

	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	song.SongID = uint(id)

	updatedSong, err := h.songsUC.Update(c, &song)
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updatedSong)
}

// Delete godoc
// @Summary Delete song
// @Description Delete by id song handler
// @Tags Song
// @Accept json
// @Produce json
// @Param id path int true "song_id"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} httpErrors.RestError "Invalid request body"
// @Failure 500 {object} httpErrors.RestError "Internal server error"
// @Router /songs/{id} [delete]
func (h *SongHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	err = h.songsUC.Delete(c, uint(id))
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// SongsByFields godoc
// @Summary Search by fields
// @Description Search songs by fields
// @Tags Song
// @Accept json
// @Produce json
// @Param title query string false "song title"
// @Param group query string false "song group"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Success 200 {object} model.SongsList
// @Router /songs/search [get]
func (h *SongHandler) SongsByFields(c *gin.Context) {
	pq, err := utils.GetPaginationFromCtx(c)
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	songsList, err := h.songsUC.SongsByFields(c, c.Query("title"), c.Query("group"), pq)
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, songsList)
}

// GetSongText godoc
// @Summary GetSongText gets song's verses with pagination
// @Description GetSongText by id song handler
// @Tags Song
// @Accept json
// @Produce json
// @Param id path int true "song_id"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Success 200 {string} string	"ok"
// @Router /songs/{id}/text [get]
func (h *SongHandler) GetSongText(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	pq, err := utils.GetPaginationFromCtx(c)
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	VersesList, err := h.songsUC.GetSongText(c, uint(id), pq)
	if err != nil {
		utils.LogResponseError(h.logger, err)
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, VersesList)
}
