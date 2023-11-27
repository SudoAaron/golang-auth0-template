package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// Assuming you have a way to access your session store
// This could be a global variable or better, part of a struct
// that includes your middleware methods.
var store *sessions.CookieStore

func init() {
	// Initialize your store, for example:
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
}

// IsAuthenticated checks if the user has already been authenticated previously.
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if user is authenticated
		if session.Values["profile"] == nil {
			session.Values["original_url"] = r.URL.String()
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// User is authenticated, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
