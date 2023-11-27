package callback

import (
	"net/http"
	"os"

	"github.com/SudoAaron/golang-auth0-template/platform/authenticator"

	"github.com/gorilla/sessions"
)

// Assuming you have a global session store
var store *sessions.CookieStore

func init() {
	// Initialize your session store
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
}

// Handler for our callback.
func Handler(auth *authenticator.Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check state parameter against session
		if r.URL.Query().Get("state") != session.Values["state"] {
			http.Error(w, "Invalid state parameter.", http.StatusBadRequest)
			return
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(r.Context(), r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to convert an authorization code into a token.", http.StatusUnauthorized)
			return
		}

		idToken, err := auth.VerifyIDToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Failed to verify ID Token.", http.StatusInternalServerError)
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set session values and save
		session.Values["access_token"] = token.AccessToken
		session.Values["profile"] = profile
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retrieve the original URL from the session
		originalURL := session.Values["original_url"]
		if originalURL == nil {
			originalURL = "/user" // Fallback URL if the original URL isn't set
		}

		// Redirect to logged in page
		// http.Redirect(w, r, "/user", http.StatusTemporaryRedirect)
		http.Redirect(w, r, originalURL.(string), http.StatusTemporaryRedirect)
	}
}
