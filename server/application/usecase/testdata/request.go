package testdata

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func BlankRequestWithSampleJwt() *http.Request {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func BookmarkGetRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/bookmarks", nil)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func BookmarkAddRequest() *http.Request {
	j, _ := json.Marshal(PbBookmark1)
	body := bytes.NewBuffer(j)
	req, _ := http.NewRequest("POST", "/bookmarks", body)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func BookmarkDeleteRequest() *http.Request {
	j, _ := json.Marshal(PbBookmark1)
	body := bytes.NewBuffer(j)
	req, _ := http.NewRequest("POST", "/bookmark/delete", body)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func BookmarkDeleteInvalidUserRequest() *http.Request {
	j, _ := json.Marshal(PbBookmark2)
	body := bytes.NewBuffer(j)
	req, _ := http.NewRequest("POST", "/bookmark/delete", body)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func OgpGetRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/ogps/"+Stat1.FQDN, nil)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func StatGetRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/stats", nil)
	return req
}

func UserIdentifyMeRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/users/me", nil)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func UserCreateRequest() *http.Request {
	j, _ := json.Marshal(PbUser1)
	body := bytes.NewBuffer(j)
	req, _ := http.NewRequest("POST", "/users", body)
	req.Header.Set("JWT", "sample jwt")
	return req
}

func UserDeleteRequest() *http.Request {
	req, _ := http.NewRequest("POST", "/users/delete", nil)
	req.Header.Set("JWT", "sample jwt")
	return req
}
