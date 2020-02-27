package ogps

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"github.com/otiai10/opengraph"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
	"time"
)

func Preview(r *http.Request) (*models.Ogp, error) {
	ogp_info, err := extractOgpInfo(r)
	return ogp_info, err
}

func Add(r *http.Request) (*models.Ogp, error) {
	ogp_info, err := extractOgpInfo(r)
	if err != nil {
		return nil, err
	}

	// confirm user
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return nil, err
	}
	if !user.IsEditor {
		return nil, errors.New("this user is not editor")
	}

	// insert
	ogp_info.UserID = user.ID
	ogp_info.CreatedAt = time.Now()
	err = ogp_info.Insert(context.Background(), db.Connection, boil.Blacklist())
	return ogp_info, err

}

func extractOgpInfo(r *http.Request) (*models.Ogp, error) {
	ogp_request := OgpRequest{}
	err := render.DecodeJSON(r.Body, &ogp_request)
	if err != nil {
		return nil, err
	}

	// 受け取った時点でFQDNであることを確認
	_, err = isValidFQDN(ogp_request.FQDN)
	if err != nil {
		return nil, err
	}

	ogp := models.Ogp{
		URL: fmt.Sprintf("https://%s", ogp_request.FQDN),
	}

	isExist, err := models.Ogps(models.OgpWhere.URL.EQ(ogp.URL)).Exists(context.Background(), db.Connection)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errors.New("ogp exist")
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

	// copy body for later use as below parse method destruct res.body
	body_copy := CopyBodyNonDestructive(res)
	ogp_info := opengraph.New(ogp.URL)
	if err := ogp_info.Parse(res.Body); err != nil {
		return nil, err
	}
	ogp_info.ToAbsURL()

	if len(ogp_info.Image) == 0 {
		return nil, errors.New("no image")
	}

	decoder, err := GetDecoder(body_copy, res.Header)
	if err != nil {
		return nil, err
	}
	title, _ := decoder.String(ogp_info.Title)
	description, _ := decoder.String(ogp_info.Description)

	new_ogp := models.Ogp{
		Type:        ogp_info.Type,
		URL:         ogp.URL,
		Title:       title,
		Description: description,
		Image:       ogp_info.Image[0].URL,
	}

	_, err = isValidOgpImageUrl(new_ogp.Image)
	if err != nil {
		return nil, err
	}

	return &new_ogp, err
}
