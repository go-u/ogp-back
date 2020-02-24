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
	"strings"
	"time"
)

func Add(r *http.Request) (*models.Ogp, error) {
	ogp_request := OgpRequest{}
	err := render.DecodeJSON(r.Body, &ogp_request)
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

	// 受け取った時点でFQDNであることを確認
	_, err = isValidFQDN(ogp_request.FQDN)
	if err != nil {
		return nil, err
	}
	ogp := models.Ogp{
		URL: fmt.Sprintf("https://%s", ogp_request.FQDN),
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

	// confirm ogp image url
	_, err = isValidOgpImageUrl(new_ogp.Image)
	if err != nil {
		return nil, err
	}

	err = new_ogp.Insert(context.Background(), db.Connection, boil.Blacklist())
	return &new_ogp, err

}

func Preview(r *http.Request) (*models.Ogp, error) {
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

	_, err = isValidOgpImageUrl(new_ogp.Image)
	if err != nil {
		return nil, err
	}

	return &new_ogp, err
}

func isValidFQDN(s string) (bool, error) {
	if strings.HasPrefix(s, "http") {
		return false, errors.New("no prefix allowed")
	}
	if strings.HasPrefix(s, "/") {
		return false, errors.New("no suffix allowed")
	}
	return true, nil
}

func isValidOgpImageUrl(s string) (bool, error) {
	if !strings.HasPrefix(s, "https") {
		return false, errors.New("image url is not https ")
	}
	return true, nil
}
