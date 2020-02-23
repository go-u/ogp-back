package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"server/api/v1/ogps"
)

func OgpRouter() http.Handler {
	r := chi.NewRouter()
	// user名が不明でも使える。プライバシー情報を含んだ詳細な情報を返す
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := ogps.Get(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := ogps.Add(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	// create
	r.Post("/preview", func(w http.ResponseWriter, r *http.Request) {
		result, err := ogps.Preview(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	return r
}
