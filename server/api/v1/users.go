package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"server/api/v1/users"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()
	// user名が不明でも使える。プライバシー情報を含んだ詳細な情報を返す
	r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		result, err := users.GetSelfDetail(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	r.Get("/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		result, err := users.GetUserDetail(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	// create
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := users.Register(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	// update
	r.Patch("/self", func(w http.ResponseWriter, r *http.Request) {
		result, err := users.Update(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, result)
		}
	})
	r.Patch("/self/avatar", func(w http.ResponseWriter, r *http.Request) {
		err := users.UpdateAvatar(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, nil)
		}
	})
	r.Delete("/self", func(w http.ResponseWriter, r *http.Request) {
		err := users.DeleteAcount(r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.Respond(w, r, nil)
		}
	})
	return r
}
