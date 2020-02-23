package ogps

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-chi/render"
	"github.com/otiai10/opengraph"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
	"time"
)

func Add(r *http.Request) (*models.Ogp, error) {
	ogp := models.Ogp{}
	err := render.DecodeJSON(r.Body, &ogp)
	if err != nil {
		return nil, err
	}
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return nil, err
	}
	if !user.IsEditor {
		return nil, errors.New("this user is not editor")
	}

	ogp.UserID = user.ID

	buf := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", ogp.URL, buf)
	if err != nil {
		//log.Fatalln(1001, err)
		return nil, err
	}

	const ua = "twitterbot"
	req.Header.Add("User-Agent", ua)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		//log.Fatalln(1002, err)
		return nil, err
	}

	ogp_info := opengraph.New(ogp.URL)
	if err := ogp_info.Parse(res.Body); err != nil {
		//log.Fatalln(1003, err)
		return nil, err
	}

	now := time.Now()
	new_ogp := models.Ogp{
		UserID:      user.ID,
		CreatedAt:   now,
		UpdatedAt:   now,
		Type:        ogp_info.Type,
		URL:         ogp.URL,
		Title:       ogp_info.Title,
		Description: ogp_info.Description,
		Image:       ogp_info.Image[0].URL,
	}

	err = new_ogp.Insert(context.Background(), db.Connection, boil.Blacklist())
	return &new_ogp, err

}

func Preview(r *http.Request) (*models.Ogp, error) {
	ogp := models.Ogp{}
	err := render.DecodeJSON(r.Body, &ogp)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", ogp.URL, buf)
	if err != nil {
		return nil, err
	}

	const ua = "twitterbot"
	req.Header.Add("User-Agent", ua)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	ogp_info := opengraph.New(ogp.URL)
	if err := ogp_info.Parse(res.Body); err != nil {
		return nil, err
	}

	if len(ogp_info.Image) == 0 {
		return nil, errors.New("no image")
	}

	new_ogp := models.Ogp{
		Type:        ogp_info.Type,
		URL:         ogp.URL,
		Title:       ogp_info.Title,
		Description: ogp_info.Description,
		Image:       ogp_info.Image[0].URL,
	}

	return &new_ogp, err
}
