package tools

import (
	"context"
	"firebase.google.com/go/auth"
	"log"
	"net/http"
	"server/initialize/firebaseauth"
)

func GetSiginProviderFromJWT(r *http.Request) (*string, error) {
	// バインディング https://qiita.com/tantakan/items/2dc435a3d20474f26e96
	verified_token, err := verifyIDToken(context.Background(), r.Header.Get("JWT"))
	if err != nil {
		return nil, err
	}
	sign_in_provider := verified_token.Claims["firebase"].(map[string]interface{})["sign_in_provider"].(string)
	return &sign_in_provider, nil
}

func GetUidFromJWT(r *http.Request) (*string, error) {
	verified_token, err := verifyIDToken(context.Background(), r.Header.Get("JWT"))
	if err != nil {
		return nil, err
	}
	return &verified_token.UID, nil
}

func verifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := firebaseauth.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}
	return token, nil
}
