package server

import (
	"EffectiveMobile_Project/config"
	"EffectiveMobile_Project/internal/entities"
	"context"
	"github.com/go-chi/chi/v5"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log/slog"
	"net/http"
)

type SongProvider interface {
	GetAllSongsFiltered(ctx context.Context, req *entities.AllSongsRequest) (*[]entities.AllSongsResponse, error)
	GetLyricsPaginated(ctx context.Context, req *entities.LyricsRequest) (*[]string, error)
	DeleteSong(ctx context.Context, songId int) error
	ChangeSongData(ctx context.Context, req *entities.ChangeSongReq) (*entities.AddSong, error)
	AddSong(ctx context.Context, req *entities.AddSong) (*entities.AddSong, error)
}

type Server struct {
	db     SongProvider
	router *chi.Mux
	log    *slog.Logger
}

func NewServer(db SongProvider, router *chi.Mux, log *slog.Logger) *Server {
	return &Server{
		db:     db,
		router: router,
		log:    log,
	}
}

func (s *Server) Run(config *config.Config) {
	srv := http.Server{
		Addr:    config.Address,
		Handler: s.router,
	}

	s.router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	s.router.Get("/all", s.GetAllSongsFiltered)
	s.router.Get("/lyrics", s.GetLyricsPaginated)
	s.router.Delete("/delete/{id}", s.DeleteSong)
	s.router.Patch("/change/{id}", s.ChangeSongData)

	if err := srv.ListenAndServe(); err != nil {
		s.log.Error("failed to start server")
	}
}
