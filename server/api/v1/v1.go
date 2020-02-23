package v1

import (
	"github.com/go-chi/chi"
	"net/http"
)

func ApiRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/users", UserRouter())
	r.Mount("/ogps", OgpRouter())
	r.Mount("/bookmarks", BookmarkRouter())
	return r
}
