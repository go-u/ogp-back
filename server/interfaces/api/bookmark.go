package api

import (
	"github.com/go-chi/render"
	"net/http"
	pb "server/etc/protocol"
)

func (h *Handler) handleGetBookmarks(w http.ResponseWriter, r *http.Request) {
	user, err := h.requestUser(r)
	if err != nil {
		err = render.Render(w, r, ErrUnAuthorized(err))
		return
	}
	bookmarks, err := h.BookmarkUsecase.Get(r.Context(), user.Id)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, bookmarks)
}

func (h *Handler) handleAddBookmark(w http.ResponseWriter, r *http.Request) {
	user, err := h.requestUser(r)
	if err != nil {
		err = render.Render(w, r, ErrUnAuthorized(err))
		return
	}

	var bookmark pb.Bookmark
	err = render.DecodeJSON(r.Body, &bookmark)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}

	err = h.BookmarkUsecase.Add(r.Context(), user.Id, bookmark.Fqdn)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, "ok")
}

func (h *Handler) handleDeleteBookmark(w http.ResponseWriter, r *http.Request) {
	user, err := h.requestUser(r)
	if err != nil {
		err = render.Render(w, r, ErrUnAuthorized(err))
		return
	}

	//bookmark := &pb.Bookmark{}
	var bookmark pb.Bookmark
	err = render.DecodeJSON(r.Body, &bookmark)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}

	err = h.BookmarkUsecase.Delete(r.Context(), user.Id, bookmark.Fqdn)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, "ok")
}
