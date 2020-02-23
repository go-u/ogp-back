package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"server/api/v1/bookmarks"
)

func BookmarkRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := bookmarks.GetBookmarks(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := bookmarks.Add(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	r.Post("/delete", func(w http.ResponseWriter, r *http.Request) {
		err := bookmarks.Delete(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, nil)
		}
	})
	return r
}
