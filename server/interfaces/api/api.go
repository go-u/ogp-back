package api

import (
	"github.com/go-chi/chi"
	"net/http"
	"server/application/usecase"
	pb "server/etc/protocol"
)

type Config struct {
	AuthUsecase     usecase.AuthUsecase
	BookmarkUsecase usecase.BookmarkUsecase
	OgpUsecase      usecase.OgpUsecase
	StatUsecase     usecase.StatUsecase
	UserUsecase     usecase.UserUsecase
}

type Handler struct {
	*Config
	router chi.Router
}

// net.httpのハンドラを満たすため必用
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewHandler(config *Config) *Handler {
	h := &Handler{
		Config: config,
		router: chi.NewRouter(),
	}

	// users
	h.router.Get("/users/me", h.handleMe)
	h.router.Post("/users", h.handleCreateAccount)
	h.router.Post("/users/delete", h.handleDeleteAccount)

	// stats
	h.router.Get("/stats", h.handleGetStats)

	// ogps
	h.router.Get("/ogps/{fqdn}", h.handleGetOgps)

	// bookmarks
	h.router.Get("/bookmarks", h.handleGetBookmarks)
	h.router.Post("/bookmarks", h.handleAddBookmark)
	h.router.Post("/bookmark/delete", h.handleDeleteBookmark)

	return h
}

func (h *Handler) requestUser(r *http.Request) (*pb.User, error) {
	jwt := r.Header.Get("JWT")
	uid, err := h.AuthUsecase.Verify(jwt)
	if err != nil {
		return nil, err
	}
	return h.UserUsecase.GetByUID(r.Context(), uid)
}
