package protected

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// Assuming you have a global session store
var store *sessions.CookieStore

// Assuming you have a global template variable
var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("web/template/*"))
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
}

// Handler for the protected page.
func Handler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile, ok := session.Values["profile"]
	if !ok {
		// Handle case where the profile is not set in the session
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = templates.ExecuteTemplate(w, "protected.html", profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
