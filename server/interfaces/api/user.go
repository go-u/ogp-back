package api

import (
	"github.com/go-chi/render"
	"net/http"
	pb "server/etc/protocol"
)

func (h *Handler) handleMe(w http.ResponseWriter, r *http.Request) {
	user, err := h.requestUser(r)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, user)
}

func (h *Handler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header.Get("JWT")
	uid, err := h.AuthUsecase.Verify(jwt)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}

	var newUser *pb.User
	err = render.DecodeJSON(r.Body, &newUser)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}

	err = h.UserUsecase.CreateAccount(r.Context(), uid, newUser.Name)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, "ok")
}

func (h *Handler) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	user, err := h.requestUser(r)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	err = h.UserUsecase.DeleteAccount(r.Context(), user.Uid)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, "ok")
}
