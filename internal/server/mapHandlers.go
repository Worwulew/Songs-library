package server

import (
	docs "github.com/Worwulew/Songs-library/docs"
	"github.com/Worwulew/Songs-library/internal/middleware"
	httpSongs "github.com/Worwulew/Songs-library/internal/songs/delivery/http"
	songsRepository "github.com/Worwulew/Songs-library/internal/songs/repository"
	"github.com/Worwulew/Songs-library/internal/songs/usecase"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// MapHandlers Map Server Handlers and inits other structs
func (s *Server) MapHandlers() (*gin.Engine, error) {
	sRepo := songsRepository.NewSongsRepository(s.db)

	docs.SwaggerInfo.BasePath = "/"

	songsUC := usecase.NewSongsUseCase(s.cfg, sRepo, s.logger)
	songHandler := httpSongs.NewSongHandlers(s.cfg, songsUC, s.logger)

	router := gin.New()

	mw := middleware.NewMiddlewareManager(s.cfg, s.logger)
	router.Use(mw.RequestLoggerMiddleware())
	router.Use(mw.DebugMiddleware())

	songs := router.Group("/songs")
	{
		songs.POST("", songHandler.Create)
		songs.PUT("/:id", songHandler.Update)
		songs.DELETE("/:id", songHandler.Delete)
		songs.GET("/search", songHandler.SongsByFields)
		songs.GET("/:id/text", songHandler.GetSongText)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router, nil
}
