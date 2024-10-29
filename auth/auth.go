package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/sessions"
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
	key := os.Getenv("SESSION_KEY")                        // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                                   // 30 days
	isProd, err := strconv.ParseBool(os.Getenv("IS_PROD")) // Set to true when serving over https
	if err != nil {
		isProd = false
	}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(GoogleOauthConfig.ClientId, GoogleOauthConfig.ClientSecret, GoogleOauthConfig.RedirectUrl, GoogleOauthConfig.Scopes...),
	)
}

func OauthGoogleCallback(res http.ResponseWriter, req *http.Request) {
	// Add provider as a query parameter (this is what `gothic.CompleteUserAuth` expects)
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	jsonOut, err := json.MarshalIndent(user, "", "")
	fmt.Println(string(jsonOut[:]))

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}

func OauthGoogleLogin(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func OauthGoogleLogout(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
}
