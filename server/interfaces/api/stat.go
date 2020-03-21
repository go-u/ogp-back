package api

import (
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) handleGetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.StatUsecase.Get(r.Context())
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, stats)
}
