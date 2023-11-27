package login

import (
	"crypto/rand"
	"encoding/base64"
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

// Handler for our login.
func Handler(auth *authenticator.Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := generateRandomState()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Save the state inside the session.
		session, err := store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["state"] = state
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the authorization URL
		http.Redirect(w, r, auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
