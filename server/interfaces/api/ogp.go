package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) handleGetOgps(w http.ResponseWriter, r *http.Request) {
	fqdn := chi.URLParam(r, "fqdn")
	tweets, err := h.OgpUsecase.Get(r.Context(), fqdn)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, tweets)
}
