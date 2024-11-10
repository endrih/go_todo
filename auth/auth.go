package auth

import (
	"context"
	"endrih/go_todo/application"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type OAuthConfig struct {
	RedirectUrl  string
	ClientId     string
	ClientSecret string
	Scopes       []string
}

var GoogleOauthConfig = OAuthConfig{
	RedirectUrl:  "http://localhost:10500/auth/google/callback",
	ClientId:     os.Getenv("GOOGLE_OAUTH2_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET"),
	Scopes: []string{
		"profile",
		"email",
	},
}

func NewAuth() {

	gothic.Store = application.App.Session

	goth.UseProviders(
		google.New(GoogleOauthConfig.ClientId, GoogleOauthConfig.ClientSecret, GoogleOauthConfig.RedirectUrl, GoogleOauthConfig.Scopes...),
	)
}

func OauthGoogleCallback(res http.ResponseWriter, req *http.Request) {
	// Add provider as a query parameter (this is what `gothic.CompleteUserAuth` expects)
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		application.App.ErrorLog.Println(err)
		return
	}
	//jsonOut, err := json.MarshalIndent(user, "", "")
	//fmt.Println(string(jsonOut[:]))
	http.Redirect(res, req, fmt.Sprintf("/login?id_token=%s", user.IDToken), http.StatusTemporaryRedirect)
}

func OauthGoogleLogin(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func OauthGoogleLogout(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
}

func verifyIdToken(idTokenInput string) *idtoken.Payload {
	payload, err := idtoken.Validate(context.Background(), idTokenInput, application.App.Config.GoogleConfig.GOOGLE_OAUTH2_CLIENT_ID)
	if err != nil {
		panic(err)
	}
	return payload
}
